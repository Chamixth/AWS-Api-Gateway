package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetResource(apiId, resourceId,region string) (*apigateway.Resource,error){
	client, err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	output, err := client.GetResource(&apigateway.GetResourceInput{
		ResourceId: aws.String(resourceId),
		RestApiId: aws.String(apiId),
	})

	if err != nil{
		return nil,err
	}

	return output,nil
}