package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"

	"github.com/gofiber/fiber/v2"
)

func GetDocumentationParts(c *fiber.Ctx) error {
	var region dto.Region
	if err := c.BodyParser(&region); err != nil{
		return err
	}
	restApiId := c.Query("restApiId")

	result, err := dao.GetDocumentationParts(restApiId, region.Region)

	if err != nil{
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
