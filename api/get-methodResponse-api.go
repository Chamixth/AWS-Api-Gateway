package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetMethodResponse(c *fiber.Ctx) error {
	var region dto.Region

	if err := c.BodyParser(&region); err != nil {
		return err
	}

	restApiId := c.Query("restApiId")
	resourceId := c.Query("resourceId")
	httpMethod := c.Query("httpMethod")
	statusCode := c.Query("statusCode")

	result, err := dao.GetMethodResponse(restApiId, resourceId, httpMethod, statusCode, region.Region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
