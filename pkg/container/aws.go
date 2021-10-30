package container

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"os"
)

// NewAWSConfig create a new AWS config object
func (c *Container) NewAWSConfig() error {
	var err error

	// Create resolver for localstack if AWS_ENDPOINT and AWS_REGION are specified
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		awsEndpoint := os.Getenv("AWS_ENDPOINT")
		awsRegion := os.Getenv("AWS_REGION")

		if awsEndpoint != "" && awsRegion != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	c.awsConfig, err = config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolver(customResolver))
	if err != nil {
		return err
	}
	c.Logger.Debug("AWS Config loaded")
	return nil
}