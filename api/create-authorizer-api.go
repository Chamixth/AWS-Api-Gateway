package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateAuthorizer(c *fiber.Ctx) error {
	var application dto.CreateAuthorizer

	if err := c.BodyParser(&application); err != nil {
		return err
	}

	result, err := dao.CreateAuthorizer(application)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
