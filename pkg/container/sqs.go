package container

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// NewSQSClient creates a new NewSQSClient client
func (c *Container) NewSQSClient() {
	if c.SQSClient == nil {
		c.SQSClient = sqs.NewFromConfig(c.AwsConfig)
	}
}
