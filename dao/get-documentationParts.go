package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetDocumentationParts(restapiId,region string) (*apigateway.GetDocumentationPartsOutput, error){
	client, err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result, err := client.GetDocumentationParts(&apigateway.GetDocumentationPartsInput{
		RestApiId: &restapiId,
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}
