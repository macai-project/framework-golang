package container

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awsV1 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

// NewAWSConfig create a new AWS config object
func (c *Container) NewAWSConfig(ctx context.Context) error {
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

	c.AwsConfig, err = config.LoadDefaultConfig(ctx, config.WithEndpointResolver(customResolver))
	if err != nil {
		return err
	}
	c.Logger.Debug("AWS Config loaded")
	return nil
}

func (c *Container) NewAWSConfigV1() error {

	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	awsRegion := os.Getenv("AWS_REGION")
	profile := os.Getenv("AWS_PROFILE")

	if awsEndpoint != "" && awsRegion != "" && profile != "" {
		c.Session = session.Must(session.NewSessionWithOptions(session.Options{
			Profile: profile,
		}))

		c.AwsConfigV1 = awsV1.Config{
			Endpoint: awsV1.String(awsEndpoint),
			Region:   awsV1.String(awsRegion),
		}
		return nil
	}

	c.Session = session.Must(session.NewSession())
	c.AwsConfigV1 = awsV1.Config{}
	c.Logger.Debug("AWS Config V1 loaded")
	return nil
}
