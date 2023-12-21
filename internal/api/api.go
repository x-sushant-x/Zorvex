/*
	PoF - This file contain all the code for providing RESTful API for client.
*/

package api

import (
	"encoding/json"
	"net/http"

	"github.com/sushant102004/zorvex/internal/agent"
	"github.com/sushant102004/zorvex/internal/types"
)

type HTTPHandler struct {
	agent agent.Agent
}

func NewHTTPHandler(svcAgent agent.Agent) *HTTPHandler {
	return &HTTPHandler{
		agent: svcAgent,
	}
}

func (h *HTTPHandler) handleRegisterService(w http.ResponseWriter, r *http.Request) {
	var serviceData types.Service

	if err := json.NewDecoder(r.Body).Decode(&serviceData); err != nil {
		WriteResponse(w, http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
		return
	}

	if err := h.agent.RegisterService(serviceData); err != nil {
		WriteResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "unable to add service to database: ",
			"error":   err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("service registered successfully"))
}

func (h *HTTPHandler) handleDiscoverService(w http.ResponseWriter, r *http.Request) {
	serviceName := r.URL.Query().Get("name")
	if serviceName == "" {
		WriteResponse(w, http.StatusBadRequest, map[string]string{
			"message": "service name must be defined in query params",
		})
		return
	}

	data, err := h.agent.GetServicesData(serviceName)

	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "unable to get services data: ",
			"error":   err.Error(),
		})
		return
	}

	if len(data) == 0 {
		WriteResponse(w, http.StatusBadRequest, map[string]string{
			"message": "service not found",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.handleRegisterService(w, r)
	case "GET":
		h.handleDiscoverService(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}
