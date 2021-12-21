package container

import (
	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

// NewAppSyncClient creates a new NewAppSyncClient client
func (c *Container) NewAppSyncClient() {
	if c.AppSyncClient == nil {
		c.AppSyncClient = appsync.NewFromConfig(c.AwsConfig)
	}
}
