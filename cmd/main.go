package main

import (
	"github.com/fitness-tracker/api"
	"github.com/fitness-tracker/internal/app"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	server := fiber.New()
	// Create a new service instance, purpose of service is to handle business logic
	service := app.NewService()

	// Create a new handler instance, purpose of handler is to handle request and response
	handler := api.NewHandler(service)

	// Define routes
	server.Get("/", handler.HelloWorld)

	err := server.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
