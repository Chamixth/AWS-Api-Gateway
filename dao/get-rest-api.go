package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetRestApi(restApiId, region string) (*apigateway.RestApi, error) {
	client, err := utils.AwsSession(region)
	if err != nil {
		return nil, err
	}

	result, err := client.GetRestApi(&apigateway.GetRestApiInput{
		RestApiId: aws.String(restApiId),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
