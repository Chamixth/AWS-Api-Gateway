package dto

type Integration struct{
	HttpMethod *string `json:"httpMethod"`
	ResourceId *string `json:"resourceId"`
	RestApiId *string `json:"restApiId"`
	CacheKeyParameters []*string `json:"cacheKeyParameters"`
	ContentHandling *string `json:"contentHandling"`
	Type *string `json:"type"`
	IntegrationType string `json:"integrationType"`
	Uri *string `json:"uri"`
	RequestParameters map[string]*string `json:"requestParameters"`
	Region string `json:"region"`
}

type IntegrationResponse struct{
	HttpMethod *string `json:"httpMethod"`
	ResourceId *string `json:"resourceId"`
	RestApiId *string `json:"restApiId"`
	StatusCode *string `json:"statusCode"`
	ContentHandling *string `json:"contentHandling"`
	ResponseParameters map[string]*string `json:"responseParameters"`
	ResponseTemplates map[string]*string `json:"responseTemplates"`
	SelectionPattern *string `json:"selectionPattern"`
	Region string `json:"region"`
}