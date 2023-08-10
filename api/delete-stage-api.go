package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteStage(c *fiber.Ctx) error {
	var region dto.Region
	if err := c.BodyParser(&region); err != nil {
		return err
	}
	restApiId := c.Query("restApiId")
	stageName := c.Query("stageName")

	err := dao.DeleteStage(restApiId, stageName, region.Region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Stage": "Deleted Successfully",
	})
}
