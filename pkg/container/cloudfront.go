package container

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

// NewCloudfrontClient create a new Cloudfront client
func (c *Container) NewCloudfrontClient() {
	if c.CloudfrontClient == nil {
		c.CloudfrontClient = cloudfront.NewFromConfig(c.AwsConfig)
	}
}
