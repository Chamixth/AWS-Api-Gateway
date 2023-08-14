package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAuthorizer(c *fiber.Ctx) error {
	var region dto.Region

	if err := c.BodyParser(&region); err != nil {
		return err
	}

	restApiId := c.Query("restApiId")
	authorizerId := c.Query("authorizerId")

	result, err := dao.GetAuthorizer(authorizerId, restApiId, region.Region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
