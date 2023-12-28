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
	workers  chan struct{}
}

func NewHealthChecker(agent agent.Agent, dbClient db.DBClient, totalWorkers int) *HealthChecker {
	return &HealthChecker{
		agent:    agent,
		dbClient: dbClient,
		workers:  make(chan struct{}, totalWorkers),
	}
}

func (h HealthChecker) StartHealthChecker() {
	services, err := h.agent.GetAllServices()
	if err != nil {
		log.Error().Msgf("Health checker error: %v", err.Error())
		log.Error().Msgf("Will try again")
		return
	}

	for _, service := range services {
		h.workers <- struct{}{}

		go func(service types.Service) {
			defer func() {
				<-h.workers
			}()
			h.CheckHealth(service)
		}(service)
	}
}

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
		log.Error().Msgf("Health check failed")
		err := h.dbClient.ChangeServiceStatus(service.ID, "down")
		if err != nil {
			log.Error().Msgf("Unable to change service status: %v", err.Error())
		}
		return
	} else {
		log.Info().Msgf("Health check passed by service: %v", service.Name)
	}
}
