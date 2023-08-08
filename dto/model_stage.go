package dto

type Stage struct{
	RestApiId string 	`json:"restApiId"`
	DeploymentId string `json:"deploymentId"`
	StageName string `json:"stageName"`
	Region string `json:"region"` 
}

