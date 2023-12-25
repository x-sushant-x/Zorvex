package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sushant102004/zorvex/internal/agent"
	"github.com/sushant102004/zorvex/internal/api"
	"github.com/sushant102004/zorvex/internal/db"
	loadbalancer "github.com/sushant102004/zorvex/internal/load-balancer"
	"github.com/sushant102004/zorvex/internal/observer"
)

func init() {
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

	observer := observer.NewObserver(db)
	log.Info().Msgf("Ready to observe services")
	observer.SetupAllServicesOnStart()

	lb := loadbalancer.NewLoadBalancer(*observer)

	agent, err := agent.NewServiceAgent(lb, db)
	if err != nil {
		log.Fatal().Msgf("unable to create new agent: %v", err.Error())
	}

	go observer.StreamInstances()

	go func() {
		handler := api.NewHTTPHandler(agent)
		log.Info().Msgf("API Handlers Running")
		handler.ServeHandlers()
	}()

	go func() {
		handler := api.NewClientHTTPHandler(agent)
		log.Info().Msgf("Client Handlers Running")
		handler.ServeHandlers()
	}()

	select {}
}
