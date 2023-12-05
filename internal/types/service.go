/*
	PoF - This file contains all the types for service registration.
*/

package types

import "time"

type Service struct {
	// ID should be unique for each service.
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	IPAddress           string    `json:"ip_address"`
	Port                string    `json:"port"`
	CreationTime        time.Time `json:"creation_time"`
	LastSyncTime        time.Time `json:"last_sync_time"`
	HealthURL           string    `json:"health_url"`
	HealthStatus        string    `json:"health_status"`         // active, unknown, down
	Endpoint            string    `json:"endpoint"`              // This is the endpoint that will be used by client to call a microservice.
	LoadBalancingMethod string    `json:"load_balancing_method"` // RoundRobin, LeastConnections, Resource, FixedWeighting
}
