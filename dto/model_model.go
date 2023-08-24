package dto

type ModelRequest struct{
	RestApiId string `json:"restApiId"`
	ContentType string `json:"contentType"`
	Description string `json:"description"`
	Name string `json:"name"`
	Schema string `json:"schema"`
	Region string `json:"region"`
}