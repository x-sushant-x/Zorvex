package agent

import (
	"fmt"

	"github.com/sushant102004/zorvex/internal/db"
	loadbalancer "github.com/sushant102004/zorvex/internal/load-balancer"
	"github.com/sushant102004/zorvex/internal/types"
)

type Agent interface {
	RegisterService(types.Service) error
	// Get all the services and send them to load balancer
	GetServicesData(string) ([]types.Service, error)
}

type ServiceAgent struct {
	db db.DBClient
	lb *loadbalancer.Balancer
}

func NewServiceAgent(lb *loadbalancer.LoadBalancer, db db.DBClient) (*ServiceAgent, error) {
	return &ServiceAgent{
		db: db,
		lb: &lb.Balancer,
	}, nil
}

func (sa *ServiceAgent) RegisterService(data types.Service) error {
	if err := sa.db.AddNewServiceToDB(data); err != nil {
		return fmt.Errorf("unable to add service to db: %v", err.Error())
	}
	return nil
}

func (sa *ServiceAgent) GetServicesData(name string) ([]types.Service, error) {
	svcInstances, err := sa.db.GetServiceInstances(name)
	if err != nil {
		return nil, fmt.Errorf("unable to get service instances: %v", err.Error())
	}
	return svcInstances, nil
}
