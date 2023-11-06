package main

import (
	handler "rest_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", handler.Hello)
}
