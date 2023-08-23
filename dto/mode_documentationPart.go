package dto

import "github.com/aws/aws-sdk-go/service/apigateway"

type CreateDocumentationPart struct {
	RestApiID string `json:"restApiId"`
	Location apigateway.DocumentationPartLocation `json:"location"`
	Properties string `json:"properties"`
	Region string `json:"region"`
}