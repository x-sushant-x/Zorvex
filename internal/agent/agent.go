package agent

import (
	loadbalancer "github.com/sushant102004/zorvex/internal/load-balancer"
	"github.com/sushant102004/zorvex/internal/types"
)

type Agent interface {
	RegisterService(types.Service) error

	// Get all the services and send them to load balancer
	GetServicesData() ([]types.Service, error)
}

type ServiceAgent struct {
	lb *loadbalancer.Balancer
}

func (sa *ServiceAgent) RegisterService(types.Service) error {
	return nil
}

func (sa *ServiceAgent) GetServicesData() ([]types.Service, error) {
	return nil, nil
}
