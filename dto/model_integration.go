package dto

type Integration struct{
	HttpMethod *string `json:"httpMethod"`
	ResourceId *string `json:"resourceId"`
	RestApiId *string `json:"restApiId"`
	Type *string `json:"type"`
	IntegrationType string `json:"integrationType"`
	Uri *string `json:"uri"`
	Region string `json:"region"`
}