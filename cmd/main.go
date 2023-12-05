package main

import (
	z_aws "github.com/sushant102004/Zorvex/internal/aws"
	"github.com/sushant102004/Zorvex/internal/db"
)

func main() {
	aws := z_aws.NewAWSSession()
	dynamoClient := db.NewDynamoDBClient(aws.Session)

	dynamoClient.AutoMigrate()
}
