package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func PutIntegration(application dto.Integration)(*apigateway.Integration,error){
	client, err := utils.AwsSession(application.Region)

	if err != nil{
		return nil,err
	} 

	result, err := client.PutIntegration(&apigateway.PutIntegrationInput{
		HttpMethod: application.HttpMethod,
		ResourceId: application.ResourceId,
		RestApiId: application.RestApiId,
		Type: application.Type,
		IntegrationHttpMethod: &application.IntegrationType,
		Uri: application.Uri,
		
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}