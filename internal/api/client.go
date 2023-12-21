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
			WriteResponse(w, http.StatusBadRequest, map[string]string{
				"message": "service name must be defined in query parameters",
			})
			return
		}

		redirectURL, err := h.agent.ServeClient(r.URL.Query().Get("service"))
		if err != nil {
			WriteResponse(w, http.StatusInternalServerError, map[string]string{
				"message": "unable to serve client",
				"error":   err.Error(),
			})
			return
		}

		if redirectURL == "" {
			WriteResponse(w, http.StatusInternalServerError, map[string]string{
				"message": "invalid redirect URL",
			})
			return
		}

		if !strings.HasPrefix(redirectURL, "http://") && !strings.HasPrefix(redirectURL, "https://") {
			redirectURL = "http://" + redirectURL
		}

		http.Redirect(w, r, redirectURL, http.StatusFound)
	default:
		WriteResponse(w, http.StatusBadRequest, map[string]string{
			"message": "method not allowed",
		})
	}
}
