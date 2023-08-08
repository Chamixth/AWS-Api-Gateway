package dto



type Method struct {
	HTTPMethod            string           `json:"httpMethod"`
	ResourceId            string           `json:"resourceId"`
	RestApiId             string           `json:"restApiId"`
	AuthorizationType     string           `json:"authorizationType"`
	RequestParameters     map[string]*bool `json:"requestParameters"`
	
	Region                string `json:"region"`
}

type MethodResponse struct {
	HTTPMethod    string             `json:"httpMethod"`
	ResourceId    string             `json:"resourceId"`
	RestApiId     string             `json:"restApiId"`
	StatusCode    string             `json:"statusCode"`
	ResponseModel map[string]*string `json:"responseParameters"`
	Region   string             `json:"region"`
}
