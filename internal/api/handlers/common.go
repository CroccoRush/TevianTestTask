package handlers

import (
	"TevianTestTask/internal/api/models"
	"github.com/gofiber/fiber/v2"
)

// Ping godoc
// @Summary      Ping
// @Description  Returns "pong" if the server is healthy
// @Tags         healthcheck
// @Produce      json
// @Success      200     {object}   models.ResponseCommon
// @Router       /api/ping          [post]
func Ping(c *fiber.Ctx) (err error) {
	var res models.ResponseCommon
	res.Create("pong")
	return c.Status(fiber.StatusOK).JSON(res)
}
