package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func CreateDocumentationPart(application dto.CreateDocumentationPart)(*apigateway.DocumentationPart,error){
	client, err := utils.AwsSession(application.Region)

	if err != nil{
		return nil,err
	}

	result, err := client.CreateDocumentationPart(&apigateway.CreateDocumentationPartInput{
		RestApiId: &application.RestApiID,
		Location: &application.Location,
		Properties: &application.Properties,
	})
	if err != nil{
		return nil, err
	}

	return result,nil
}