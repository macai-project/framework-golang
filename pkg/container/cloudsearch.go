package container

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
)

// NewCloudsearchClient create a new Cloudsearch client
func (c *Container) NewCloudsearchClient(endpoint string) {
	if c.CloudsearchClient == nil {
		c.CloudsearchClient = cloudsearchdomain.NewFromConfig(
			c.AwsConfig,
			cloudsearchdomain.WithEndpointResolver(
				cloudsearchdomain.EndpointResolverFromURL(
					endpoint,
				),
			),
		)
	}
}
