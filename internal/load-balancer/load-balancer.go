package loadbalancer

import "fmt"

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
	fmt.Println("Round Robin Strategy")
	return nil
}

func (lb *LoadBalancer) Random() error {
	fmt.Println("Random Strategy")
	return nil
}

func (lb *LoadBalancer) LeastConnections() error {
	fmt.Println("Least Connection Strategy")
	return nil
}

func (lb *LoadBalancer) LeastResponseTime() error {
	fmt.Println("Least Response Strategy")
	return nil
}
