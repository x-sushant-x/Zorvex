/*
	PoF - This file contains all the types for service registration.
*/

package types

type Service struct {
	ID                  string       `json:"id" rethinkdb:"id, omitempty"`
	Name                string       `json:"name" rethinkdb:"name"`
	Tags                []string     `json:"tags" rethinkdb:"tags"`
	HTTPMethod          string       `json:"http_method" rethinkdb:"http_method"`
	IPAddress           string       `json:"ip_address" rethinkdb:"ip_address"`
	Port                int          `json:"port" rethinkdb:"port"`
	RegisterTime        string       `json:"register_time" rethinkdb:"register_time"`
	LastSyncTime        string       `json:"last_sync_time" rethinkdb:"last_sync_time"`
	Endpoint            string       `json:"endpoint" rethinkdb:"endpoint"`                           // This is the endpoint that will be used by client to call a microservice.
	LoadBalancingMethod string       `json:"load_balancing_method" rethinkdb:"load_balancing_method"` // RoundRobin, LeastConnections, Resource, FixedWeighting
	TotalConnections    int          `json:"total_connections" rethinkdb:"total_connections"`
	DeRegisterAfter     int          `json:"de_register_after" rethinkdb:"de_register_after"`
	Status              string       `json:"status" rethinkdb:"status"` // active, unknown, down
	HealthConfig        HealthConfig `json:"health_config"`
}

type HealthConfig struct {
	HealthCheckEndpoint string              `json:"health_check_endpoint" rethinkdb:"health_check_endpoint"` // <ip_address>:<port>/health
	Interval            int                 `json:"interval" rethinkdb:"interval"`
	Options             HealthConfigOptions `json:"options" rethinkdb:"options"`
}

type HTTPHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type HealthConfigOptions struct {
	Headers            []HTTPHeader `json:"http_headers" rethinkdb:"http_headers"`
	Body               any          `json:"body" rethinkdb:"body"` // This must be encoded into json.
	ExpectedStatusCode int          `json:"expected_status_code" rethinkdb:"expected_status_code"`
}
