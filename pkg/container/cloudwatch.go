package container

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

// NewCloudwatchClient create a new Cloudwatch client
func (c *Container) NewCloudwatchClient() {
	if c.CloudwatchClient == nil {
		c.CloudwatchClient = cloudwatch.NewFromConfig(c.AwsConfig)
	}
}
