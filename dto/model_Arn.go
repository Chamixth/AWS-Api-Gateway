package dto

type MethodRequestArn struct{
	MethodRequestArn string `json:"methodRequestArn"`
}

type MethodRequestArnRequest struct{
	RestApiId string `json:"restApiId"`
	HttpMethod string `json:"httpMethod"`
	AccountId string `json:"accountId"`
	Region string `json:"region"`
}