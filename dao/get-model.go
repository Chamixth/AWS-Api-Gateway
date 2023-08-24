package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetModel(restApiID, modelName, region string) (*apigateway.Model,error){
	client, err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result, err := client.GetModel(&apigateway.GetModelInput{
		RestApiId: &restApiID,
		ModelName: &modelName,
	})
	return result,nil
}