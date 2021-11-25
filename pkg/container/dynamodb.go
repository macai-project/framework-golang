package container

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamoDBV1 "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/guregu/dynamo"
)

// NewDynamoDBClient create a new DynamoDBClient client
func (c *Container) NewDynamoDBClient() {
	if c.DynamoDBClient == nil {
		c.DynamoDBClient = dynamodb.NewFromConfig(c.awsConfig)
	}
}

// NewDynamoDBORM create a new DynamoDBORM client
func (c *Container) NewDynamoDBORM() {
	if c.DynamoDBORM == nil {
		c.DynamoDBORM = dynamo.New(c.Session, &c.awsConfigV1)
	}
}

// NewDynamoDBORM create a new DynamoDBORM client
func (c *Container) NewDynamoDBORMXray() {
	if c.DynamoDBORM == nil {
		customClient := dynamoDBV1.New(c.Session, &c.awsConfigV1)
		xray.AWS(customClient.Client)
		c.DynamoDBORM = dynamo.NewFromIface(customClient)
	}
}