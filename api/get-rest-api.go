package api

import (
	"aws-api-gateway/dao"
	"aws-api-gateway/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetRestApi(c *fiber.Ctx) error {
	var region dto.Region
	apiId := c.Query("apiId")

	if err := c.BodyParser(&region); err != nil {
		return err
	}

	 result, err := dao.GetRestApi(apiId, region.Region) 
	 if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(result)
}
