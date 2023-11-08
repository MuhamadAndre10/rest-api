package main

import (
	"context"
	"errors"
	"log"
	"rest_api/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
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

	// make a connection server with ngrok connection
	run(context.Background(), app)

}

func run(ctx context.Context, app *fiber.App) error {
	tun, err := ngrok.Listen(ctx, config.HTTPEndpoint(),
		ngrok.WithAuthtoken("2XNKNQ4oQJBde3uIuThENrt5HGA_37gtbJMecgeEMfAUzG9JW"), ngrok.WithRegion("ap"))

	var nerr ngrok.Error

	if errors.As(err, &nerr) {
		log.Fatal("Something went wrong:,", nerr.Msg())
		return err
	}

	defer tun.Close()

	log.Println("tunnel created:", tun.URL())

	return app.Listener(tun)
}
