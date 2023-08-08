package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func PutMethodResponse(c *fiber.Ctx) error {
	var application dto.MethodResponse

	if err := c.BodyParser(&application); err != nil {
		return err
	}

	result, err := dao.PutMethodResponse(application)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
