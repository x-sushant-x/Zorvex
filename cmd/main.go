package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sushant102004/zorvex/internal/agent"
	"github.com/sushant102004/zorvex/internal/api"
	"github.com/sushant102004/zorvex/internal/db"
	healtchecker "github.com/sushant102004/zorvex/internal/health_checker"
	loadbalancer "github.com/sushant102004/zorvex/internal/load-balancer"
	"github.com/sushant102004/zorvex/internal/observer"
)

func init() {
	// Initializing Zerolog for better logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123})
}

func main() {
	// Create database connections
	db, err := db.NewRethinkClient()
	log.Info().Msgf("Connected to database")

	if err != nil {
		log.Fatal().Err(err).Msgf("unable to create database connection")
	}

	// Onserver is something that store services data in memory for better load balancer performance.
	observer := observer.NewObserver(db)
	log.Info().Msgf("Ready to observe services")

	// If there are already some services in the database than this function will automatically add them to memory.
	observer.SetupAllServicesOnStart()

	// Streaming service data changes in real time
	go observer.StreamInstances()

	// Instance of load balancer
	lb := loadbalancer.NewLoadBalancer(*observer)

	// Agent is the central men who puts all the things together
	agent, err := agent.NewServiceAgent(lb, db)
	if err != nil {
		log.Fatal().Msgf("unable to create new agent: %v", err.Error())
	}

	// APIs for agent
	go func() {
		handler := api.NewHTTPHandler(agent)
		log.Info().Msgf("API Handlers Running")
		handler.ServeHandlers()
	}()

	// APIs for client
	go func() {
		handler := api.NewClientHTTPHandler(agent)
		log.Info().Msgf("Client Handlers Running")
		handler.ServeHandlers()
	}()

	// Health checker for checking health of services
	healthChecker := healtchecker.NewHealthChecker(agent, db, 10)

	go func() {
		for {
			// Change this time according to your needs
			time.Sleep(time.Minute * 30)
			healthChecker.StartHealthChecker()
		}
	}()

	healthChecker.StartHealthChecker()

	select {}
}
