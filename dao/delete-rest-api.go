package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func DeleteRestApi(restApiId, region string) error {
	client, err := utils.AwsSession(region)
	if err != nil {
		return err
	}
	if _, err := client.DeleteRestApi(&apigateway.DeleteRestApiInput{
		RestApiId: aws.String(restApiId),
	}); err != nil {
		return err
	}
	return nil
}
