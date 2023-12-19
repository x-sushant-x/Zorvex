package api

import (
	"net/http"
	"strings"

	"github.com/sushant102004/zorvex/internal/agent"
)

type ClientHTTPHandler struct {
	agent agent.Agent
}

func NewClientHTTPHandler(agent agent.Agent) *ClientHTTPHandler {
	return &ClientHTTPHandler{
		agent: agent,
	}
}

func (h *ClientHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Query().Get("service") == "" {
			http.Error(w, "service name must be defined in query parameters", http.StatusBadRequest)
			return
		}

		redirectURL, err := h.agent.ServeClient(r.URL.Query().Get("service"))
		if err != nil {
			http.Error(w, "unable to serve client: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if redirectURL == "" {
			http.Error(w, "invalid redirect URL", http.StatusInternalServerError)
			return
		}

		if !strings.HasPrefix(redirectURL, "http://") && !strings.HasPrefix(redirectURL, "https://") {
			redirectURL = "http://" + redirectURL
		}

		http.Redirect(w, r, redirectURL, http.StatusFound)
	default:
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}
