package container

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

// NewLambdaClient creates a new NewLambdaClient client
func (c *Container) NewLambdaClient() {
	if c.LambdaClient == nil {
		c.LambdaClient = lambda.NewFromConfig(c.AwsConfig)
	}
}
