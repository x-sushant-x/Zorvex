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
	agent agent.ServiceAgent
}

func NewHTTPHandler(svcAgent *agent.ServiceAgent) *HTTPHandler {
	return &HTTPHandler{
		agent: *svcAgent,
	}
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var serviceData types.Service

		if err := json.NewDecoder(r.Body).Decode(&serviceData); err != nil {
			http.Error(w, "invalid request body"+err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.agent.RegisterService(serviceData); err != nil {
			http.Error(w, "unable to add service to database: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("service registered successfully"))
	default:
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}
