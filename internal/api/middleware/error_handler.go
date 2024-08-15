package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func PanicErrorHandler(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			err = c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}()
	return c.Next()
}
