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

	// make a connection server.
	log.Fatal(app.Listen(":3000"))

}
