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
	app.Post("/createStage",api.CreateStage)
	app.Put("/putMethodResponse",api.PutMethodResponse)
	app.Get("/getMethod",api.GetMethod)
	app.Put("/putIntegrationResponse",api.PutIntegrationResponse)
	app.Get("/getStage",api.GetStage)
	app.Get("/getIntegration",api.GetIntegration)
	app.Get("/getMethodRequestArn",api.GetMethodRequestArn)
	app.Delete("/deleteMethod",api.DeleteMethod)
	app.Delete("/deleteStage",api.DeleteStage)
	app.Delete("/deleteIntegration",api.DeleteIntegration)
	app.Get("/getResources",api.GetResources)
	app.Get("/getRestApis",api.GetRestApis)
	app.Get("/getMethodResponse",api.GetMethodResponse)
	app.Get("/getIntegrationResponse",api.GetIntegrationResponse)
	app.Get("/getDeployment",api.GetDeployment)
	app.Get("/getAuthorizer",api.GetAuthorizer)
	app.Post("/createAuthorizer",api.CreateAuthorizer)
	app.Delete("/deleteAuthorizer",api.DeleteAuthorizer)
	app.Get("/getAuthorizers",api.GetAuthorizers)
	app.Get("/getDocumentationParts",api.GetDocumentationParts)
	app.Post("/createDocumentationPart",api.CreateDocumentationPart)
	app.Post("/getExport",api.GetExport)
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
