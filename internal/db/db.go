package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBClient struct {
	db *dynamodb.DynamoDB
}

func NewDynamoDBClient(sess *session.Session) *DynamoDBClient {
	db := dynamodb.New(sess)
	return &DynamoDBClient{db: db}
}

func (d *DynamoDBClient) AutoMigrate() {
	// AutoMigrate function will setup all the required
	err := d.CreateServiceTable()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DynamoDBClient) CreateServiceTable() error {
	tableName := "services"

	output, err := d.db.ListTables(&dynamodb.ListTablesInput{
		ExclusiveStartTableName: &tableName,
	})
	if err != nil {
		return err
	}

	if len(output.TableNames) == 0 {
		input := &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("Name"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("IPAddress"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("Port"),
					AttributeType: aws.String("N"),
				},
				{
					AttributeName: aws.String("CreationTime"),
					AttributeType: aws.String("N"),
				},
				{
					AttributeName: aws.String("LastSyncTime"),
					AttributeType: aws.String("N"),
				},
				{
					AttributeName: aws.String("HealthURL"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("HealthStatus"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("Endpoint"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("LoadBalancingMethod"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("ID"),
					KeyType:       aws.String("S"),
				},
			},
		}

		_, err := d.db.CreateTable(input)
		if err != nil {
			return err
		}

		fmt.Println("Table created successfully.")

	} else {
		fmt.Println("Table already created!")
	}

	return nil
}
