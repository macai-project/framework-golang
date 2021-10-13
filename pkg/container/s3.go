package container

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewS3Client create a new S3 client
func (c *Container) NewS3Client() {
	c.S3Client = s3.NewFromConfig(c.awsConfig)
}
