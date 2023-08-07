package dto

type Method struct{
	HTTPMethod string `json:"httpMethod"`
	ResourceId string `json:"resourceId"`
	RestApiId string `json:"restApiId`
	AuthorizationType string `json:"authorizationType"`
	Region string `json:"region"`
}