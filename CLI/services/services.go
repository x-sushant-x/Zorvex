package services

import (
	"encoding/json"
	"net/http"

	types "github.com/sushant102004/Zorvex/CLI/types"
	"github.com/sushant102004/Zorvex/CLI/utils"
)

type ServiceResponse struct {
	Services []types.Service `json:"services"`
}

func GetAllServices() ([]types.Service, error) {
	resp, err := http.Get("http://localhost:3000/all-services")
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
