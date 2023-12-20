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
)

type Balancer interface {
	RoundRobin([]types.Service) string
}

type LoadBalancer struct {
	ob *observer.Observer

	// Round Robin stuff
	rrMux *sync.Mutex
}

func NewLoadBalancer(ob observer.Observer) *LoadBalancer {
	return &LoadBalancer{
		ob:    &ob,
		rrMux: &sync.Mutex{},
	}
}

func (balancer *LoadBalancer) checkAlive(url string) bool {
	conn, err := net.DialTimeout("tcp", url, time.Second*1)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func (lb *LoadBalancer) RoundRobin(service string) (string, error) {
	lb.rrMux.Lock()
	defer lb.rrMux.Unlock()

	services := lb.ob.ServicesInstances[service]

	if len(services) == 0 {
		return "", errors.New("no service found with name: " + service)
	}

	idx := lb.ob.ServicesPointers[service]

	for {
		targetService := services[idx]
		urlStr := fmt.Sprintf("http://%s:%d", targetService.IPAddress, targetService.Port)

		isAlive := lb.checkAlive(targetService.IPAddress + ":" + fmt.Sprint(targetService.Port))

		log.Info().Msgf("Target Index: %d", idx)

		if isAlive {
			lb.ob.ServicesPointers[service] = (idx + 1) % len(services)
			return urlStr, nil
		}

		log.Error().Msgf("Service not alive. Proceeding to next service.")

		idx = (idx + 1) % len(services)

		if idx == lb.ob.ServicesPointers[service] {
			// If we have gone through all instances and reached the starting point, exit the loop
			break
		}
	}

	return "", errors.New("all instances of the service are not alive")
}
