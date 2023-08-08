package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetStage(c *fiber.Ctx) error {
	var region dto.Region
	restApiId := c.Query("restApiId")
	stageName := c.Query("stageName")

	if err := c.BodyParser(&region); err != nil {
		return err
	}

	result, err := dao.GetStage(restApiId, stageName, region.Region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
