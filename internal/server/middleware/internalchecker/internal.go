package internalchecker

import "github.com/gofiber/fiber/v2"

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		if headers["X-Forwarded-For"] != "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}
