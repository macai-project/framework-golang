package container

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
)

// NewAWSConfig create a new AWS config object
func (c *Container) NewAWSConfig() error {
	var err error
	c.awsConfig, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	c.Logger.Debug("AWS Config loaded")
	return nil
}