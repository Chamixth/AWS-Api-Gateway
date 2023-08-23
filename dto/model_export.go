package dto

type GetExport struct{
	Accepts string `json:"accepts"`
	ExportType string `json:"exportType"`
	RestApiId string `json:"restApiId"`
	StageName string `json:"stageName"`
	Region string `json:"region"`
}