package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"

	"github.com/gofiber/fiber/v2"
)

func GetModel(c *fiber.Ctx) error {
	var region dto.Region
	if err := c.BodyParser(&region); err != nil {
		return err
	}
	restApiID := c.Query("restApiId")
	modelName := c.Query("modelName")
	result,err := dao.GetModel(restApiID, modelName, region.Region)

	if err != nil{
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
