package dao

import (
	"aws-api-gateway/utils"

	"github.com/aws/aws-sdk-go/service/apigateway"
)

func GetDeployment(restApiId, deploymentId, region string) (*apigateway.Deployment,error){
	client,err := utils.AwsSession(region)

	if err != nil{
		return nil,err
	}

	result,err := client.GetDeployment(&apigateway.GetDeploymentInput{
		RestApiId: &restApiId,
		DeploymentId: &deploymentId,
	})

	if err != nil{
		return nil,err
	}
	return result,nil
}