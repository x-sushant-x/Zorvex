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
	GetServicesData() ([]types.Service, error)
}

type ServiceAgent struct {
	db db.RethinkClient
	lb *loadbalancer.Balancer
}

func NewServiceAgent(lb *loadbalancer.LoadBalancer, db *db.RethinkClient) (*ServiceAgent, error) {
	return &ServiceAgent{
		db: *db,
		lb: &lb.Balancer,
	}, nil
}

func (sa *ServiceAgent) RegisterService(data types.Service) error {
	if err := sa.db.AddNewServiceToDB(data); err != nil {
		return fmt.Errorf("unable to add service to db: %v", err.Error())
	}
	return nil
}

func (sa *ServiceAgent) GetServicesData() ([]types.Service, error) {
	return nil, nil
}
