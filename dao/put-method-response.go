package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func PutMethodResponse(application dto.MethodResponse)(*apigateway.MethodResponse,error){
	client, err := utils.AwsSession(application.Region)

	if err != nil{
		return nil,err
	}

	result, err := client.PutMethodResponse(&apigateway.PutMethodResponseInput{
		HttpMethod: &application.HTTPMethod,
		ResourceId: &application.ResourceId,
		RestApiId: &application.RestApiId,
		StatusCode: &application.StatusCode,
		ResponseModels: application.ResponseModel,
		
	})

	if err != nil{
		return nil,err
	}

	return result,err
}