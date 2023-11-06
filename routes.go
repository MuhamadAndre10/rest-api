package main

import (
	handler "rest_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", handler.Hello)
	app.Get("/products", handler.FetchAllProducts)
	app.Get("/products/:id", handler.FetchProduct)

	app.Post("/products", handler.InsertProduct)

	app.Delete("/products/:id", handler.DeleteProduct)

}
