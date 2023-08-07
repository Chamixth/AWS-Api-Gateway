package main

import (
	"aws-api-gateway/api"
	"log"

	"github.com/gofiber/fiber/v2"

)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 40 * 1024 * 1024, // Set the body limit to 40MB
	})

	app.Post("/createRestApi", api.CreateRestApi)
	app.Delete("/deleteRestApi",api.DeleteRestApi)
	app.Get("/getRestApi",api.GetRestApi)
	app.Get("/getResource",api.GetResource)
	app.Post("/createResource",api.CreateResource)
	app.Delete("/deleteResource",api.DeleteResource)
	app.Put("/putMethod",api.PutMethod)
	app.Put("putIntegration",api.PutIntegration)
	app.Post("/createDeployment",api.CreateDeployment)
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
