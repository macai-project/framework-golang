package container

import (
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

// SESClient create a new SES client
func (c *Container) NewSESClient() {
	if c.SESClient == nil {
		c.SESClient = ses.NewFromConfig(c.AwsConfig)
	}
}
