package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func DeleteResource(c *fiber.Ctx) error {
	var region dto.Region
	restApiId := c.Query("apiId")
	resourceId := c.Query("resourceId")

	if err := c.BodyParser(&region); err != nil {
		return err
	}

	err := dao.DeleteResource(resourceId, restApiId, region.Region)

	if err != nil {
		return nil
	}

	return c.Status(http.StatusOK).JSON("Resource Successfully Deleted")
}
