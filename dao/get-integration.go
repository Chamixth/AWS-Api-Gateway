package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetIntegration(httpMethod, resourceId, restApiId, region string) (*apigateway.Integration,error){
	client,err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result, err := client.GetIntegration(&apigateway.GetIntegrationInput{
		RestApiId: &restApiId,
		ResourceId: &resourceId,
		HttpMethod: &httpMethod,
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}