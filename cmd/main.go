package main

import (
	"fmt"
	"net/http"

	"github.com/sushant102004/zorvex/internal/agent"
	"github.com/sushant102004/zorvex/internal/api"
	"github.com/sushant102004/zorvex/internal/db"
	loadbalancer "github.com/sushant102004/zorvex/internal/load-balancer"
)

func main() {
	// Create database connections
	db, err := db.NewRethinkClient()

	if err != nil {
		panic(err)
	}

	lb := loadbalancer.NewLoadBalancer()

	agent, err := agent.NewServiceAgent(lb, db)
	if err != nil {
		fmt.Println("unable to create agent")
		panic(err)
	}

	handler := api.NewHTTPHandler(agent)

	http.ListenAndServe(":3000", handler)

	select {}
}
