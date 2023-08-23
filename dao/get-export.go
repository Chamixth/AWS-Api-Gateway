package dao

import (
	"aws-api-gateway/dto"
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetExport(application dto.GetExport)(*apigateway.GetExportOutput,error){
	client, err := utils.AwsSession(application.Region)

	if err != nil{
		return nil,err
	}

	result, err := client.GetExport(&apigateway.GetExportInput{
		Accepts: &application.Accepts,
		ExportType: &application.ExportType,
		RestApiId: &application.RestApiId,
		StageName: &application.StageName,
	})

	if err != nil{
		return nil,err
	}

	return result,nil
}