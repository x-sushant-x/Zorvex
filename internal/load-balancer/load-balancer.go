package loadbalancer

type Balancer interface {
	RoundRobin() error // Read about it later
	Random() error
	LeastConnections() error
	LeastResponseTime() error // Read about it later
}

type LoadBalancer struct{ Balancer }

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{}
}

func (lb *LoadBalancer) RoundRobin() error {
	return nil
}

func (lb *LoadBalancer) Random() error {
	return nil
}

func (lb *LoadBalancer) LeastConnections() error {
	return nil
}

func (lb *LoadBalancer) LeastResponseTime() error {
	return nil
}
