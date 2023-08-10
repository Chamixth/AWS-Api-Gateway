package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetMethodResponse(restApiId, resourceId,httpMethod,statusCode,region string) (*apigateway.MethodResponse,error){
	client,err := utils.AwsSession(region)

	if err != nil{
		return nil, err
	}

	result,err := client.GetMethodResponse(&apigateway.GetMethodResponseInput{
		ResourceId: &resourceId,
		RestApiId: &restApiId,
		HttpMethod: &httpMethod,
		StatusCode: &statusCode,
	})

	if err != nil{
		return nil,err
	}

	return result,err
}