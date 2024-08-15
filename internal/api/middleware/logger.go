package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func Logger(c *fiber.Ctx) error {
	start := time.Now()
	defer log.Printf(
		"%s %s %s",
		c.Method(),
		c.BaseURL(),
		time.Since(start),
	)
	return c.Next()
}
