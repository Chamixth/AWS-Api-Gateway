package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func CreateRestApi(application dto.RestApi) (*string, error) {
	client, err := utils.AwsSession(application.Region)

	if err != nil {
		fmt.Println("Error creating session:", err)
		return nil, err
	}
	input, err := client.CreateRestApi(&apigateway.CreateRestApiInput{
		Name: aws.String(application.Name),
		Description: aws.String(application.Description),
	})

	if err != nil {
		fmt.Println("Error creating REST API:", err)
		return nil,err
	}

	return input.Id,err
}
