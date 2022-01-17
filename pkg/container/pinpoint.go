package container

import (
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
)

// NewPinpointClient creates a new NewPinpointClient client
func (c *Container) NewPinpointClient() {
	if c.PinpointClient == nil {
		c.PinpointClient = pinpoint.NewFromConfig(c.AwsConfig)
	}
}
