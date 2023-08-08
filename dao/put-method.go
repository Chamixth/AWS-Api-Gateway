package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func PutMethod(application dto.Method) (*apigateway.Method,error) {
	client, err := utils.AwsSession(application.Region) 
	if err != nil{
		return nil,err
	}

	result,err := client.PutMethod(&apigateway.PutMethodInput{
		HttpMethod: aws.String(application.HTTPMethod),
		ResourceId: aws.String(application.ResourceId),
		RestApiId: aws.String(application.RestApiId),
		AuthorizationType: aws.String(application.AuthorizationType),
		RequestParameters:application.RequestParameters ,
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}