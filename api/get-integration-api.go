package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetIntegration(c *fiber.Ctx) error {
	var region dto.Region
	restApiId := c.Query("restApiId")
	resourceId := c.Query("resourceId")
	httpMethod := c.Query("httpMethod")

	if err := c.BodyParser(&region); err != nil {
		return err
	}

	result, err := dao.GetIntegration(httpMethod, resourceId, restApiId, region.Region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
