package api

import (
	"TevianTestTask/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc fiber.Handler
}

type Routes []Route

func setupRoutes() {
	for _, route := range routes {
		app.Add(route.Method, route.Pattern, route.HandlerFunc)
	}
}

var routes = Routes{
	Route{
		"Ping",
		strings.ToUpper("Post"),
		"/api/ping",
		handlers.Ping,
	},

	Route{
		"AddTask",
		strings.ToUpper("Post"),
		"/api/task/add",
		handlers.AddTask,
	},

	Route{
		"GetTask",
		strings.ToUpper("Get"),
		"/api/task",
		handlers.GetTask,
	},

	Route{
		"ProcessTask",
		strings.ToUpper("Post"),
		"/api/task/process",
		handlers.ProcessTask,
	},

	Route{
		"DeleteTask",
		strings.ToUpper("Delete"),
		"/api/task",
		handlers.DeleteTask,
	},

	Route{
		"UploadImage",
		strings.ToUpper("POST"),
		"/api/image/upload",
		handlers.UploadImage,
	},
}
