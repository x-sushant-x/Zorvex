package api

import (
	"github.com/gofiber/fiber/v2"
)

func WriteResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	c.Request().Header.Add("Content-Type", "application/json")

	return c.Status(statusCode).JSON(data)
}
