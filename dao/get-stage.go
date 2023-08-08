package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetStage(restApiId, stageName,region string) (*apigateway.Stage,error){
	client,err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result, err := client.GetStage(&apigateway.GetStageInput{
		RestApiId: &restApiId,
		StageName: &stageName,
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}