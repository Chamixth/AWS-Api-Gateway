package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func DeleteAuthorizer(restApiId, authorizerId, region string) error {
	client, err := utils.AwsSession(region)

	if err != nil{
		return err
	}

	if _,err := client.DeleteAuthorizer(&apigateway.DeleteAuthorizerInput{
		RestApiId: &restApiId,
		AuthorizerId: &authorizerId,
	}); err != nil{
		return err
	}

	return nil
}