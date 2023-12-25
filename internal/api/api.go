/*
	PoF - This file contain all the code for providing RESTful API for client.
*/

package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
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

func (h *HTTPHandler) ServeHandlers() {
	app := fiber.New()

	app.Get("/discover", h.handleDiscoverService)
	app.Get("/all-services", h.handleGetAllServices)
	app.Post("/register", h.handleRegisterService)

	app.Listen(":3000")
}

func (h *HTTPHandler) handleRegisterService(c *fiber.Ctx) error {
	var serviceData types.Service

	if err := c.BodyParser(&serviceData); err != nil {
		return WriteResponse(c, http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
	}

	if err := h.agent.RegisterService(serviceData); err != nil {
		return WriteResponse(c, http.StatusInternalServerError, map[string]string{
			"message": "unable to add service to database: ",
			"error":   err.Error(),
		})

	}
	return WriteResponse(c, http.StatusOK, map[string]string{
		"message": "service registered successfully",
	})
}

func (h *HTTPHandler) handleDiscoverService(c *fiber.Ctx) error {
	serviceName := c.Query("service")
	if serviceName == "" {
		return WriteResponse(c, http.StatusBadRequest, map[string]string{
			"message": "service name must be defined in query params",
		})
	}

	data, err := h.agent.GetServiceData(serviceName)

	if err != nil {
		return WriteResponse(c, http.StatusInternalServerError, map[string]string{
			"message": "unable to get services data: ",
			"error":   err.Error(),
		})
	}

	if len(data) == 0 {
		return WriteResponse(c, http.StatusBadRequest, map[string]string{
			"message": "service not found",
		})
	}

	return WriteResponse(c, http.StatusOK, map[string]any{
		"service": data,
	})
}

func (h *HTTPHandler) handleGetAllServices(c *fiber.Ctx) error {
	services, err := h.agent.GetAllServices()
	if err != nil {
		return WriteResponse(c, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if len(services) == 0 {
		return WriteResponse(c, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return WriteResponse(c, http.StatusOK, map[string]any{
		"services": services,
	})
}
