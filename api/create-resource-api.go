package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateResource(c *fiber.Ctx) error {
	var application dto.Resource


	if err := c.BodyParser(&application); err != nil {
		return err
	}

	result, err := dao.CreateResource(application)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
