package container

import "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

// NewCognitoClient creates a new NewCognitoClient client
func (c *Container) NewCognitoClient() {
	if c.CognitoClient == nil {
		c.CognitoClient = cognitoidentityprovider.NewFromConfig(c.AwsConfig)
	}
}
