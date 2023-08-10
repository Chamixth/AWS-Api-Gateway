package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func DeleteMethod(http_method, resourceId, restApiId,region string) (*apigateway.DeleteMethodOutput,error) {
	client, err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result,err := client.DeleteMethod(&apigateway.DeleteMethodInput{
		HttpMethod: &http_method,
		ResourceId: &resourceId,
		RestApiId: &restApiId,
	})
	
	if err != nil{
		return nil,err
	}

	return result,nil
}