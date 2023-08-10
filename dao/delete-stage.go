package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func DeleteStage(restApiId, stageName, region string) error {
	client, err := utils.AwsSession(region)

	if err != nil{
		return err
	}

	if _,err := client.DeleteStage(&apigateway.DeleteStageInput{
		RestApiId: &restApiId,
		StageName: &stageName,
	}); err != nil{
		return err
	}

	return nil
}