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
