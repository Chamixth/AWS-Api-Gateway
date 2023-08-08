package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetMethod(restApiId, resourceId, httpMethod, region string) (*apigateway.Method,error){
	client,err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result,err := client.GetMethod(&apigateway.GetMethodInput{
		HttpMethod: aws.String(httpMethod),
		RestApiId: aws.String(restApiId),
		ResourceId: aws.String(resourceId),
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}