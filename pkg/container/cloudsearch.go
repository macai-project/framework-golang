package container

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
)

// NewCloudsearchClient create a new Cloudsearch client
func (c *Container) NewCloudsearchClient() {
	if c.CloudsearchClient == nil {
		c.CloudsearchClient = cloudsearch.NewFromConfig(c.AwsConfig)
	}
}
