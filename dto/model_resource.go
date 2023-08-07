package dto

type Resource struct{
	RestApiId  string `json:"restApiId"`
	ParentId string `json:"parentId"`
	Region string `json:"region"`
	PathPart string `json:"pathPart"`
}

type CreateResourceResponse struct {
	ResourceID  string `json:"resourceId"`
	ResourceURL string `json:"resourceUrl"`
}