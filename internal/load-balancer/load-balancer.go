package loadbalancer

import (
	"errors"
	"fmt"
	"sync"

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

func (lb *LoadBalancer) RoundRobin(service string) (string, error) {
	lb.rrMux.Lock()

	services := lb.ob.ServicesInstances[service]

	if len(service) == 0 {
		return "", errors.New("no service found with name: " + service)
	}

	targetIdx := (lb.ob.ServicesPointers[service] + 1) % len(services)

	lb.ob.ServicesPointers[service]++
	lb.rrMux.Unlock()

	targetService := services[targetIdx]

	urlStr := fmt.Sprintf("http://%s:%d", targetService.IPAddress, targetService.Port)

	return urlStr, nil
}
