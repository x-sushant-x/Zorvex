package main

import (
	"github.com/sushant102004/zorvex/internal/db"
)

func main() {
	// Create database connections
	_, err := db.NewRethinkClient()

	if err != nil {
		panic(err)
	}

	// Create database tables if not exists (first time application launch)

	// Start agent and listen for incoming service registration and de registration requests.

	select {}
}
