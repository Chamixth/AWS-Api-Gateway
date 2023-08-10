package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetIntegrationResponse(httpMethod, restApiId, resourceId, statusCode,region string) (*apigateway.IntegrationResponse,error){
	client,err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result,err := client.GetIntegrationResponse(&apigateway.GetIntegrationResponseInput{
		RestApiId: &restApiId,
		ResourceId: &resourceId,
		HttpMethod: &httpMethod,
		StatusCode: &statusCode,
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}