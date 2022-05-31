package container

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// NewSSMClient creates a new NewSSMClient client
func (c *Container) NewSSMClient() {
	if c.SSMClient == nil {
		c.SSMClient = ssm.NewFromConfig(c.AwsConfig)
	}
}
