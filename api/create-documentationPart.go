package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"

	"github.com/gofiber/fiber/v2"
)

func CreateDocumentationPart(c *fiber.Ctx) error {
	var application dto.CreateDocumentationPart

	if err := c.BodyParser(&application); err != nil{
		return err
	}

	result, err := dao.CreateDocumentationPart(application)

	if err != nil{
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
