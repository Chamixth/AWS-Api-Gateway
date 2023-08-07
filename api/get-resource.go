package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetResource(c *fiber.Ctx) error {
	var region dto.Region
	restApiId := c.Query("apiId")
	resourceId := c.Query("resourceId")

	if err := c.BodyParser(&region); err != nil {
		return err
	}

	result, err := dao.GetResource(restApiId, resourceId, region.Region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
