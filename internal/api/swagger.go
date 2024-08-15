package api

import "github.com/gofiber/contrib/swagger"

var configDefault = swagger.Config{
	Next:     nil,
	BasePath: "/",
	FilePath: "./internal/api/docs/swagger.json",
	Path:     "swagger",
	Title:    "TevianTestTask API documentation",
	CacheAge: 1, // Default to 1 hour
}

func useSwagger() {
	app.Use(swagger.New(configDefault))
}
