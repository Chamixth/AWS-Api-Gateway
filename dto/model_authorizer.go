package dto

type CreateAuthorizer struct {
	Name string `json:"name"`
	RestApiId string `json:"restApiId"`
	AuthorizerResultTtlInSeconds int `json:"authorizerResultTtlInSeconds"`
	Type string `json:"type"`
	IdentitySource string `json:"identitySource"`
	AuthType string `json:"authType"`
	AuthorizerUri string `json:"authorizerUri"`
	IdentityValidationExpression string `json:"identityValidationExpression"`
	ProvidersArn []*string `json:"providersArn"`
	AuthorizerCredentials string `json:"authorizerCredentials"`
	Region string `json:"region"`
}