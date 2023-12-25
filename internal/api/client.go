package api

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
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

func (h *ClientHTTPHandler) ServeHandlers() {
	app := fiber.New()

	app.Get("/", h.ServeClient)

	app.Listen(":3001")
}

func (h *ClientHTTPHandler) ServeClient(c *fiber.Ctx) error {
	serviceName := c.Query("service")

	if serviceName == "" {
		return WriteResponse(c, http.StatusBadRequest, map[string]string{
			"message": "service name must be defined in query parameters",
		})
	}

	redirectURL, err := h.agent.ServeClient(serviceName)
	if err != nil {
		return WriteResponse(c, http.StatusInternalServerError, map[string]string{
			"message": "unable to serve client",
			"error":   err.Error(),
		})
	}

	if redirectURL == "" {
		return WriteResponse(c, http.StatusInternalServerError, map[string]string{
			"message": "invalid redirect URL",
		})
	}

	if !strings.HasPrefix(redirectURL, "http://") && !strings.HasPrefix(redirectURL, "https://") {
		redirectURL = "http://" + redirectURL
	}

	return c.Redirect(redirectURL, http.StatusFound)

}
