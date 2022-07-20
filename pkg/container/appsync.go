package container

import (
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	appsync "github.com/sony/appsync-client-go"
	"github.com/sony/appsync-client-go/graphql"
)

// NewAppSyncClient creates a new AppSyncClient client
func (c *Container) NewAppSyncClient(url string, region string) {
	if c.AppSyncClient == nil {
		sess := session.Must(session.NewSession())
		signer := v4.NewSigner(sess.Config.Credentials)

		c.AppSyncClient = appsync.NewClient(
			appsync.NewGraphQLClient(graphql.NewClient(url)),
			appsync.WithIAMAuthorization(*signer, region, url),
		)
	}
}
