package dao

import (
	"aws-api-gateway/dto"
	"fmt"
)

func GetMethodRequestArn(application dto.MethodRequestArnRequest) dto.MethodRequestArn {
	methodRequestARN := fmt.Sprintf("arn:aws:execute-api:%s:%s:%s/*/%s/", application.Region, application.AccountId, application.RestApiId,application.HttpMethod)
	response := dto.MethodRequestArn{
		MethodRequestArn: methodRequestARN,
	}
	return response
}
