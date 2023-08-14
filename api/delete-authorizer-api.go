package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteAuthorizer(c *fiber.Ctx) error {
	var region dto.Region

	if err := c.BodyParser(&region); err != nil {
		return err
	}

	restApiId := c.Query("restApiId")
	authorizerId := c.Query("authorizerId")

	err := dao.DeleteAuthorizer(restApiId, authorizerId, region.Region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Authorizer": restApiId,
		"Status":     "Delete Successfully",
	})
}
