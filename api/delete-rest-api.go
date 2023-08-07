package api

import (
	"aws-api-gateway/dao"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func DeleteRestApi(c *fiber.Ctx) error {
	region := c.Query("region")
	apiId := c.Query("apiIdd")

	err := dao.DeleteRestApi(apiId, region)

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON("Rest Api Deleted Successfully")
}
