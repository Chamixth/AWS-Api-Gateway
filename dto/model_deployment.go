package dto

type Deployment struct{
	RestApiId string `json:"restApiId"`
	StageName string `json:"stageName"`
	Region string `json:"region"`
}