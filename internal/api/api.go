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
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.agent.RegisterService(serviceData); err != nil {
		http.Error(w, "unable to add service to database: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("service registered successfully"))
}

func (h *HTTPHandler) handleDiscoverService(w http.ResponseWriter, r *http.Request) {
	serviceName := r.URL.Query().Get("name")
	if serviceName == "" {
		http.Error(w, "service name must be defined in query params", http.StatusBadRequest)
		return
	}

	data, err := h.agent.GetServicesData(serviceName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "unable to get services: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		resp := map[string]string{
			"message": "service not found",
		}
		WriteResponse(w, http.StatusNoContent, resp)
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
