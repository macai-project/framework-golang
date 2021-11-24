package container

import (
	"context"
	"fmt"
	awsV1 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws/session"
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

func (c *Container) NewAWSConfigV1() error {

	c.Session = session.Must(session.NewSession())

	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	awsRegion := os.Getenv("AWS_REGION")

	if awsEndpoint != "" && awsRegion != "" {
		c.awsConfigV1 = awsV1.Config{
            Endpoint: awsV1.String(awsEndpoint),
            Region: awsV1.String(awsRegion),
        }
		c.Logger.Debug("AWS Config V1 loaded")
		return nil
	}

	return fmt.Errorf("no endpoint or region set")
}