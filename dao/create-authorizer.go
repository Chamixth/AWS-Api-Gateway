package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func CreateAuthorizer(application dto.CreateAuthorizer) (*apigateway.Authorizer, error) {
	client, err := utils.AwsSession(application.Region)

	if err != nil {
		return nil,err
	}

	result,err := client.CreateAuthorizer(&apigateway.CreateAuthorizerInput{
		RestApiId: &application.RestApiId,
		Name: &application.Name,
		AuthType: &application.AuthType,
		ProviderARNs: application.ProvidersArn,
		Type: &application.Type,
		IdentityValidationExpression: &application.IdentityValidationExpression,
		AuthorizerUri: &application.AuthorizerUri,
		IdentitySource: &application.IdentitySource,
		AuthorizerCredentials: &application.AuthorizerCredentials,
	})
	
	if err != nil{
		return nil,err
	}

	return result,nil
}
