package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetMethodRequestArn(c *fiber.Ctx) error {
	var application dto.MethodRequestArnRequest
	if err := c.BodyParser(&application); err != nil{
		return err
	}

	result := dao.GetMethodRequestArn(application)

	return c.Status(http.StatusOK).JSON(result)
}
