package container

import (
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

// NewXRayClient creates a new XRay client
func (c *Container) NewXRayClient() {

	if c.XRayClient == nil {
		c.XRayClient = xray.NewFromConfig(c.AwsConfig)
	}
}
