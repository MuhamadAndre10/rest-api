package main

import (
	"log"
	"rest_api/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// make connection to connect postgres database.
	database.ConnecDb()

	// instance a new fiber
	app := fiber.New()

	//set up middleware, if needed
	app.Use(cors.New())

	// make a route and handlers
	SetupRoutes(app)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "route not found",
		}) // => 404 "Not Found"
	})

	// make a connection server.
	log.Fatal(app.Listen(":3000"))

}
