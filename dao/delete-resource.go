package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func DeleteResource(resourceId, restApiId, region string) error {
	client, err := utils.AwsSession(region)

	if err != nil {
		return err
	}

	if _, err := client.DeleteResource(&apigateway.DeleteResourceInput{
		RestApiId:  aws.String(restApiId),
		ResourceId: aws.String(resourceId),
	}); err != nil {
		return err
	}
	return nil
}
