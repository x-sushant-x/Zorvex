package services

import (
	"encoding/json"
	"net/http"

	types "github.com/sushant102004/Zorvex/CLI/types"
	"github.com/sushant102004/Zorvex/CLI/utils"
)

var AgentURL string

func SetAgentURL(url string) {
	AgentURL = url
}

type ServiceResponse struct {
	Services []types.Service `json:"services"`
}

type InstancesResponse struct {
	Services []types.Service `json:"instances"`
}

func GetAllServices() ([]types.Service, error) {
	resp, err := http.Get(AgentURL + "/all-services")
	if err != nil {
		utils.Error(err.Error())
	}
	defer resp.Body.Close()

	var serviceResp ServiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&serviceResp); err != nil {
		utils.Error(err.Error())
	}

	return serviceResp.Services, nil
}

func GetAllDownServices() ([]types.Service, error) {
	services, err := GetAllServices()
	if err != nil {
		utils.Error(err.Error())
	}

	var downServices []types.Service
	for _, service := range services {
		if service.Status == "down" {
			downServices = append(downServices, service)
		}
	}

	return downServices, nil
}

func GetSingleService(name string) (types.Service, error) {
	resp, err := http.Get(AgentURL + "/discover?service=" + name)

	if err != nil {
		utils.Error(err.Error())
	}
	defer resp.Body.Close()

	var instancesResp InstancesResponse
	if err := json.NewDecoder(resp.Body).Decode(&instancesResp); err != nil {
		utils.Error(err.Error())
	}

	if len(instancesResp.Services) == 0 {
		return types.Service{}, nil
	}

	return instancesResp.Services[0], nil
}
