package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateRestApi(c *fiber.Ctx) error {
	var application dto.RestApi
	// Set up an AWS session
	if err := c.BodyParser(&application); err != nil {
		fmt.Println("Error with payload", err)
		return err
	}

	result, err := dao.CreateRestApi(application)

	if err != nil {
		return err
	}
	

	return c.Status(http.StatusOK).JSON(fiber.Map{"Id":result})

}
