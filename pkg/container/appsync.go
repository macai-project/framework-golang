package container

import (
	appsync "github.com/sony/appsync-client-go"
	"github.com/sony/appsync-client-go/graphql"
)

// NewAppSyncClient creates a new AppSyncClient client
func (c *Container) NewAppSyncClient(url string) {
	if c.AppSyncClient == nil {
		c.AppSyncClient = appsync.NewClient(appsync.NewGraphQLClient(graphql.NewClient(url)))
	}
}
