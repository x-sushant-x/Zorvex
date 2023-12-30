package healtchecker

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/sushant102004/zorvex/internal/agent"
	"github.com/sushant102004/zorvex/internal/db"
	"github.com/sushant102004/zorvex/internal/types"
)

type HealthChecker struct {
	agent    agent.Agent
	dbClient db.DBClient
	// Workers are the number of parallel instances running for health check.
	// Chan is used to do concurrency and limit number of workers to prevent excessive load on system
	workers chan struct{}
}

func NewHealthChecker(agent agent.Agent, dbClient db.DBClient, totalWorkers int) *HealthChecker {
	return &HealthChecker{
		agent:    agent,
		dbClient: dbClient,
		workers:  make(chan struct{}, totalWorkers),
	}
}

// Gets all the services and their instances from the database and send them to CheckHealth for individual service
func (h HealthChecker) StartHealthChecker() {
	// Calling agent.GetAllServices() that calls db.GetAllServices under the hood
	services, err := h.agent.GetAllServices()
	if err != nil {
		log.Error().Msgf("Health checker error: %v", err.Error())
		log.Error().Msgf("Will try again")
		return
	}

	for _, service := range services {
		// This is very command concurrency pattern that limits the number of goroutines creation.
		// For example is totalWorkers are 10 than only 10 goroutine functions will be created and 10 instances will be health check at that time.
		// Other instances will wait in the queue for their turn

		h.workers <- struct{}{}

		go func(service types.Service) {
			// When CheckHealth function is completed the struct that we sent to the channel will be pulled out making space for next service.
			defer func() {
				<-h.workers
			}()
			h.CheckHealth(service)
		}(service)
	}
}

// Checks health for individual service and change it's status to "down" if health check fails
func (h *HealthChecker) CheckHealth(service types.Service) {
	protocol := service.Protocol
	endpoint := service.HealthConfig.HealthCheckEndpoint
	url := fmt.Sprintf(protocol + "://" + endpoint)

	headers := service.HealthConfig.Options.Headers

	req, err := http.NewRequest(service.HTTPMethod, url, nil)
	if err != nil {
		log.Error().Msgf("Health checker error: %v", err.Error())
	}

	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != service.HealthConfig.Options.ExpectedStatusCode {
		log.Error().Msgf("Health check failed for service: %v with error: %v", service.ID, err.Error())
		err := h.dbClient.ChangeServiceStatus(service.ID, "down")
		if err != nil {
			log.Error().Msgf("Unable to change service status: %v", err.Error())
		} else {
			log.Error().Msgf("Changed status to down for service: %v", service.ID)
		}
		return
	} else {
		log.Info().Msgf("Health check passed by service: %v", service.Name)
	}
}
