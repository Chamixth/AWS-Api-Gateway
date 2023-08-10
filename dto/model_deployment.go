package dto

import "github.com/aws/aws-sdk-go/service/apigateway"

type Deployment struct {
	RestApiId string `json:"restApiId"`
	StageName string `json:"stageName"`
	Region    string `json:"region"`
}

type Deployment_Response struct {
	Deployment apigateway.Deployment `json:"deployment"`
	InvokeUrl string `json:"invokeUrl"`
}