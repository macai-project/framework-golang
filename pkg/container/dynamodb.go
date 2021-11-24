package container

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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
