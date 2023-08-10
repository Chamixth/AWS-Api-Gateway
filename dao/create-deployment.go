package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func CreateDeployment(application dto.Deployment) (*dto.Deployment_Response, error) {
	client, err := utils.AwsSession(application.Region)

	if err != nil {
		return nil, err
	}

	result, err := client.CreateDeployment(&apigateway.CreateDeploymentInput{
		RestApiId: aws.String(application.RestApiId),
		StageName: &application.StageName,
	})

	invokeURL := fmt.Sprintf("https://%s.execute-api.%s.amazonaws.com/%s", application.RestApiId, application.Region, application.StageName)

	if err != nil {
		return nil, err
	}
	response := dto.Deployment_Response{
		Deployment: *result,
		InvokeUrl:  invokeURL,
	}

	return &response, nil
}
