package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func AwsSession(region string) (*apigateway.APIGateway, error) {
	// Set up an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		// Add any other necessary configuration options here
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return nil, err
	}

	// Create an API Gateway service client
	svc := apigateway.New(sess)

	return svc,nil
}
