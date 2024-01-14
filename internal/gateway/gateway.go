/*
	It will check if the request is authenticated or not. That's it. Pretty simple.

	ToDo:
		1. Learn about vulnerabilities in this auth and CSRF.
*/

package gateway

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sushant102004/zorvex/internal/utils"
)

func CheckAuthentication(ctx *fiber.Ctx) error {
	authHeaderSecret := ctx.Get("AuthHeaderSecret")

	if authHeaderSecret == "" {
		return ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": utils.ErrUnauthorized.Error(),
		})
	}

	return ctx.Next()
}
