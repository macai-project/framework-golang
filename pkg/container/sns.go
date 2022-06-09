package container

import (
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// NewSNSClient creates a new NewSNSClient client
func (c *Container) NewSNSClient() {
	if c.SNSClient == nil {
		c.SNSClient = sns.NewFromConfig(c.AwsConfig)
	}
}
