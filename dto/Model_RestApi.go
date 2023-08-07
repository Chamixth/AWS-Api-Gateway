package dto

import "github.com/aws/aws-sdk-go/service/apigateway"

type RestApi struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Region      string `json:"region"`
}

type GetRestApi struct {
	APIKeySource              string             `json:"apiKeySource"`
	BinaryMediaTypes          []*string          `json:"binaryMediaTypes"`
	CreatedDate               int64              `json:"createdDate"`
	Description               string             `json:"description"`
	DisableExecuteApiEndpoint bool               `json:"disableExecuteApiEndpoint"`
	EndpointConfiguration     *apigateway.EndpointConfiguration       `json:"endpointConfiguration"`
	ID                        string             `json:"id"`
	MinimumCompressionSize    int                `json:"minimumCompressionSize"`
	Name                      string             `json:"name"`
	Policy                    string             `json:"policy"`
	Tags                      map[string]*string `json:"tags"`
	Version                   string             `json:"version"`
	Warning                   []*string          `json:"warnings"`
}
