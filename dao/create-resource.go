package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func CreateResource(application dto.Resource) (*dto.CreateResourceResponse, error) {
	client, err := utils.AwsSession(application.Region)

	if err != nil {
		return nil, err
	}

	result, err := client.CreateResource(&apigateway.CreateResourceInput{
		RestApiId: aws.String(application.RestApiId),
		ParentId:  aws.String(application.ParentId),
		PathPart: aws.String(application.PathPart),
	})

	if err != nil {
		return nil, err
	}

	output := dto.CreateResourceResponse{
		ResourceID: *result.Id,
		ResourceURL: *result.Path,
	}

	return &output, nil

}
