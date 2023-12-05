/*
	This file will provide client of AWS. I decided to put a seperate folder for aws in order to seperate aws configuration from core logic.
	z_aws is my custom package name for aws related files. This is done to prevent conflicts with aws golang sdk.
*/

package z_aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AWSClient struct {
	Session *session.Session
}

func NewAWSSession() *AWSClient {
	config := aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(false),
		Endpoint:                      aws.String("http://localhost:8000"),
		Region:                        aws.String("ap-south-1"),
	}

	sess, err := session.NewSession(&config)
	if err != nil {
		log.Fatal(err)
	}

	return &AWSClient{
		Session: sess,
	}
}
