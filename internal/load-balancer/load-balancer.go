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
	// TODO: Research and change timeout
	conn, err := net.DialTimeout("tcp", url, time.Second*4)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func (lb *LoadBalancer) Balance(service string) (string, error) {
	services := lb.ob.ServicesInstances[service]

	if len(services) == 0 {
		return "", errors.New("no service found with name: " + service)
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

func (lb *LoadBalancer) RoundRobin(name string, services []types.Service) (string, error) {
	lb.rrMux.Lock()
	defer lb.rrMux.Unlock()

	idx := lb.ob.ServicesPointers[name]

	for {
		targetService := services[idx]
		urlStr := fmt.Sprintf("http://%s:%d", targetService.IPAddress, targetService.Port)

		isAlive := lb.checkAlive(targetService.IPAddress + ":" + fmt.Sprint(targetService.Port))

		log.Info().Msgf("Target Index: %d", idx)

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
