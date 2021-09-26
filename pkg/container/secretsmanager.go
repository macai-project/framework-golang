package container

import (
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// NewSecretsManagerClient create a new SecretsManager client
func (c *Container) NewSecretsManagerClient() {
	c.SecretsManagerClient = secretsmanager.NewFromConfig(c.awsConfig)
}
