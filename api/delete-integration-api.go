package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteIntegration(c *fiber.Ctx) error {
	var region dto.Region
	if err := c.BodyParser(&region); err != nil {
		return err
	}
	restApiId := c.Query("restApiId")
	httpMethod := c.Query("httpMethod")
	resourceId := c.Query("resourceId")

	err := dao.DeleteIntegration(httpMethod, resourceId, restApiId, region.Region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Integration": "Delete Successfully",
	})
}
