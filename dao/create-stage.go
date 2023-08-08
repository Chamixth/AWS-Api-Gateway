package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func CreateStage(application dto.Stage)(*apigateway.Stage,error){
	client, err := utils.AwsSession(application.Region)

	if err != nil{
		return nil,err
	}

	result, err:= client.CreateStage(&apigateway.CreateStageInput{
		DeploymentId: aws.String(application.DeploymentId),
		StageName: aws.String(application.StageName),
		RestApiId: aws.String(application.RestApiId),
	})

	if err != nil{
		return nil,err
	}

	return result,err
}