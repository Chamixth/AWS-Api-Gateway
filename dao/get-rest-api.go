package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetRestApi(restApiId, region string) (*dto.GetRestApi, error) {
	client, err := utils.AwsSession(region)
	if err != nil {
		return nil, err
	}

	result, err := client.GetRestApi(&apigateway.GetRestApiInput{
		RestApiId: aws.String(restApiId),
	})
	if err != nil {
		return nil, err
	}

	output := dto.GetRestApi{}

	if result.ApiKeySource != nil {
		output.APIKeySource = *result.ApiKeySource
	}

	if result.BinaryMediaTypes != nil {
		output.BinaryMediaTypes = result.BinaryMediaTypes
	}

	if result.CreatedDate != nil {
		createdDate := time.Unix(result.CreatedDate.Unix(), 0)
		output.CreatedDate = createdDate.Unix()
	}

	if result.Description != nil {
		output.Description = *result.Description
	}

	if result.DisableExecuteApiEndpoint != nil {
		output.DisableExecuteApiEndpoint = *result.DisableExecuteApiEndpoint
	}

	if result.EndpointConfiguration != nil {
		output.EndpointConfiguration = result.EndpointConfiguration
	}

	if result.Id != nil {
		output.ID = *result.Id
	}

	if result.MinimumCompressionSize != nil {
		output.MinimumCompressionSize = int(*result.MinimumCompressionSize)
	}

	if result.Name != nil {
		output.Name = *result.Name
	}

	if result.Policy != nil {
		output.Policy = *result.Policy
	}

	if result.Tags != nil {
		output.Tags = result.Tags
	}

	if result.Version != nil {
		output.Version = *result.Version
	}

	if result.Warnings != nil {
		output.Warning = result.Warnings
	}

	return &output, nil
}
