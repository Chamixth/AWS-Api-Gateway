package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func CreateModel(application dto.ModelRequest) (*apigateway.Model, error){
	client, err := utils.AwsSession(application.Region)

	if err != nil{
		return nil, err
	}

	result, err := client.CreateModel(&apigateway.CreateModelInput{
		RestApiId: &application.RestApiId,
		ContentType: &application.ContentType,
		Description: &application.Description,
		Name: &application.Name,
		Schema: &application.Schema,

	})


	if err != nil{
		return nil,err
	}

	return result,nil
}
