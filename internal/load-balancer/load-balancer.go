package loadbalancer

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sushant102004/zorvex/internal/observer"
	"github.com/sushant102004/zorvex/internal/types"
	"github.com/sushant102004/zorvex/internal/utils"
)

// Base interface for LoadBalancer
type Balancer interface {
	RoundRobin([]types.Service) string
	// Other strategies can be implemented here
}

type LoadBalancer struct {
	ob *observer.Observer

	// Round Robin stuff
	// We need to lock the service before serving it to client and unlock it after it's work is done.
	// This is done to prevent race conditions so that no 2 clients can concurrently access the same service
	// As round robin moves to next service after serving current service
	rrMux *sync.Mutex
}

func NewLoadBalancer(ob observer.Observer) *LoadBalancer {
	return &LoadBalancer{
		ob:    &ob,
		rrMux: &sync.Mutex{},
	}
}

func (balancer *LoadBalancer) checkAlive(url string) bool {
	// TODO: Research and change timeout
	// Checking if a TCP connection can be made to the service url
	// IDK if it is needed or not. But let's keep it for more security
	conn, err := net.DialTimeout("tcp", url, time.Second*4)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// Generic Load Balancer that gets all the "active" services from observer and send it to appropriate load balancing strategy
func (lb *LoadBalancer) Balance(service string) (string, error) {
	// These services may contain the services that are down so we need to filter for "active" services before sending them to load balancer
	allServices := lb.ob.ServicesInstances[service]

	// It will hold healthy and active services
	services := []types.Service{}

	// Appending active services
	for _, s := range allServices {
		if s.Status == "active" {
			services = append(services, s)
		}
	}

	if len(services) == 0 {
		return "", errors.New("no active service found with name: " + service)
	}

	svc := services[0]

	switch svc.LoadBalancingMethod {
	case "RoundRobin":
		// This url is the serveable url
		url, err := lb.RoundRobin(service, services)
		if err != nil {
			return "", err
		}
		return url, nil
	}

	return "", utils.ErrUnableToLoadBalance
}

// Actual implementation for round robin strategy
func (lb *LoadBalancer) RoundRobin(name string, services []types.Service) (string, error) {
	lb.rrMux.Lock()
	defer lb.rrMux.Unlock()

	// Observer also store the index for the service that was last served.
	idx := lb.ob.ServicesPointers[name]

	for {
		targetService := services[idx]
		urlStr := fmt.Sprintf("http://%s:%d", targetService.IPAddress, targetService.Port)

		isAlive := lb.checkAlive(targetService.IPAddress + ":" + fmt.Sprint(targetService.Port))

		if isAlive {
			lb.ob.ServicesPointers[name] = (idx + 1) % len(services)
			return urlStr, nil
		}

		log.Error().Msgf("Service not alive. Proceeding to next service.")

		idx = (idx + 1) % len(services)

		if idx == lb.ob.ServicesPointers[name] {
			// If we have gone through all instances and reached the starting point, exit the loop
			break
		}
	}

	return "", utils.ErrNoServiceAlive
}
