package internal

import "github.com/sushant102004/Zorvex/internal/types"

type CSM interface {
	// This method will provide functionality for accepting new service registration request.
	AcceptRegistration(types.Service) error

	// This method will provide functionality to save new service into database
	RegisterService(types.Service) error
}

type CentralSystemManager struct{}

func NewCentralSystemManager() *CentralSystemManager {
	return &CentralSystemManager{}
}
