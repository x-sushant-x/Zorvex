package agent

import (
	"github.com/rs/zerolog/log"
	"github.com/sushant102004/zorvex/internal/db"
	loadbalancer "github.com/sushant102004/zorvex/internal/load-balancer"
	"github.com/sushant102004/zorvex/internal/types"
)

type Agent interface {
	// Call database function to store service data into database
	RegisterService(types.Service) error
	// Get all the instances of a service
	GetServiceData(string) ([]types.Service, error)
	// Get all the services
	GetAllServices() ([]types.Service, error)
	// Call load balancer function to get appropriate services url and redirect user to that instance
	ServeClient(string) (string, error)
}

type ServiceAgent struct {
	// Database dependency
	db db.DBClient
	// Load balancer dependency
	lb *loadbalancer.LoadBalancer
}

// These arguments are provided while calling this function (done in main.go)
func NewServiceAgent(lb *loadbalancer.LoadBalancer, db db.DBClient) (*ServiceAgent, error) {
	return &ServiceAgent{
		db: db,
		lb: lb,
	}, nil
}

func (sa *ServiceAgent) RegisterService(data types.Service) error {
	if err := sa.db.AddNewServiceToDB(data); err != nil {
		log.Err(err).Msgf("Unable to register new service: %v", err.Error())
		return err
	}
	return nil
}

func (sa *ServiceAgent) GetServiceData(name string) ([]types.Service, error) {
	svcInstances, err := sa.db.GetServiceInstances(name)
	if err != nil {
		log.Err(err).Msgf("Unable to get instance of the service: %v", err.Error())
		return nil, err
	}
	return svcInstances, nil
}

func (sa *ServiceAgent) GetAllServices() ([]types.Service, error) {
	services, err := sa.db.GetAllServices()
	if err != nil {
		log.Err(err).Msgf("Unable to get all services: %v", err.Error())
		return nil, err
	}
	return services, nil
}

func (sa *ServiceAgent) ServeClient(service string) (string, error) {
	url, err := sa.lb.Balance(service)

	if err != nil {
		return "", err
	}
	return url, nil
}
