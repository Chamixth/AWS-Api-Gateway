package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetRestApis(region string) (*apigateway.GetRestApisOutput, error) {
	client, err := utils.AwsSession(region)

	if err != nil {
		return nil, err
	}

	result, err := client.GetRestApis(&apigateway.GetRestApisInput{})

	if err != nil {
		return nil, err
	}

	return result, nil
}
