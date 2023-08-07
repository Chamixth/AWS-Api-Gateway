package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func CreateDeployment(application dto.Deployment) (*apigateway.Deployment, error) {
	client, err := utils.AwsSession(application.Region)

	if err != nil {
		return nil, err
	}

	result, err := client.CreateDeployment(&apigateway.CreateDeploymentInput{
		RestApiId: aws.String(application.RestApiId),
	})

	if err !=nil{
		return nil,err
	}

	return result,nil
}
