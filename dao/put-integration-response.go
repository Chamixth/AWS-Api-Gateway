package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func PutIntegrationResponse(application dto.IntegrationResponse) (*apigateway.IntegrationResponse, error) {
	client, err := utils.AwsSession(application.Region)

	if err != nil {
		return nil, err
	}

	result, err := client.PutIntegrationResponse(&apigateway.PutIntegrationResponseInput{
		ResourceId:         application.ResourceId,
		RestApiId:          application.RestApiId,
		HttpMethod:         application.HttpMethod,
		StatusCode:         application.StatusCode,
		ContentHandling:    application.ContentHandling,
		ResponseParameters: application.ResponseParameters,
		ResponseTemplates:  application.ResponseTemplates,
		SelectionPattern:   application.SelectionPattern,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
