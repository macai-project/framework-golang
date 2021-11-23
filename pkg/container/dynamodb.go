package container

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// NewDynamoDbClient create a new DynamoDb client
func (c *Container) NewDynamoDbClient() {
	if c.DynamoDbClient == nil {
		c.DynamoDbClient = dynamodb.NewFromConfig(c.awsConfig)
    }
}
