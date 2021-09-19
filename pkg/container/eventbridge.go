package container

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

// NewEventbridgeClient create a new Eventbridge client
func (c *Container) NewEventbridgeClient() {
	c.EventBridgeClient = eventbridge.NewFromConfig(c.AwsConfig)
}
