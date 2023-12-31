package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func DeleteMethod(http_method, resourceId, restApiId,region string) (error) {
	client, err := utils.AwsSession(region)

	if err != nil{
		return err
	}

	if _,err := client.DeleteMethod(&apigateway.DeleteMethodInput{
		HttpMethod: &http_method,
		ResourceId: &resourceId,
		RestApiId: &restApiId,
	}); 
	
	 err != nil{
		return err
	}

	return nil
}