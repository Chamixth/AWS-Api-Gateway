package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetMethod(c *fiber.Ctx) error {
	var region dto.Region

	if err := c.BodyParser(&region); err != nil {
		return err
	}

	resourceId := c.Query("resourceId")
	restApiId := c.Query("restApiId")
	httpMethod := c.Query("httpMethod")

	result, err := dao.GetMethod(restApiId, resourceId, httpMethod, region.Region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
