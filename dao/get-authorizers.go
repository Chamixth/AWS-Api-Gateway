package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetAuthorizers(restApiId,region string) (*apigateway.GetAuthorizersOutput,error){
	client,err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result, err := client.GetAuthorizers(&apigateway.GetAuthorizersInput{
		RestApiId: &restApiId,
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}