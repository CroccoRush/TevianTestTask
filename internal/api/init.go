package api

import (
	_ "TevianTestTask/internal/api/docs"
	"TevianTestTask/internal/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type fiberApp struct {
	*fiber.App
}

func (app fiberApp) Close() error {

	return app.Shutdown()
}

var (
	app fiberApp
)

// @title           TevianTestTask API documentation
// @version         1.0
// @description     This is the server for the Tevian test task.
// @contact.name    Kiselyov Vladimir
// @host      localhost:3000
// @securityDefinitions.basic  BasicAuth
func init() {

	app.App = fiber.New()

	app.Use(cors.New())
	app.Use(middleware.Logger)
	app.Use(middleware.PanicErrorHandler)

	useSwagger()

	setupRoutes()
}
