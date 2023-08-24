package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetModels(restApiId, region string) (*apigateway.GetModelsOutput,error){
	client,err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result,err := client.GetModels(&apigateway.GetModelsInput{
		RestApiId: &restApiId,
	})

	return result,nil
}