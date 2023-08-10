package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetResources(restApiId,region string) (*apigateway.GetResourcesOutput,error) {
	client, err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result, err := client.GetResources(&apigateway.GetResourcesInput{
		RestApiId: &restApiId,
	})

	if err != nil{
		return nil,err
	}

	return result,err
}