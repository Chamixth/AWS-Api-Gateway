package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func DeleteIntegration(httpMethod, resourceId, restApiId, region string) error {
	client, err := utils.AwsSession(region)

	if err != nil{
		return err
	}

	if _,err := client.DeleteIntegration(&apigateway.DeleteIntegrationInput{
		HttpMethod: &httpMethod,
		RestApiId: &restApiId,
		ResourceId: &resourceId,
	}); err != nil{
		return err
	}

	return nil
}