package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetAuthorizer(authorizerId, restApiId, region string) (*apigateway.Authorizer,error){
	client, err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result,err := client.GetAuthorizer(&apigateway.GetAuthorizerInput{
		RestApiId: &restApiId,
		AuthorizerId: &authorizerId,
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}