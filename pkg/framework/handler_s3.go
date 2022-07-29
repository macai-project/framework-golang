package framework

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/macai-project/framework-golang/pkg/container"
)

var businessLogicHandlerS3 func(ctx context.Context, c *container.Container, e events.S3Event) (string, error)

func RegisterBusinessLogicS3(f func(ctx context.Context, c *container.Container, e events.S3Event) (string, error)) {
	businessLogicHandlerS3 = f
}

func HandleRequestS3(ctx context.Context, e events.S3Event) (string, error) {
	var err error

	// Check if the container has been initialized
	if c == nil {
		return "", fmt.Errorf("container struct must be initialized and passed to the framework")
	}

	// Logger
	c.NewLogger()
	defer c.Logger.Sync()

	// AWS Config
	err = c.NewAWSConfig(ctx)
	if err != nil {
		return "error initializing AWS Config", err
	}

	// Xray
	awsv2.AWSV2Instrumentor(&c.AwsConfig.APIOptions)

	ctx, err = xray.ContextWithConfig(ctx, xray.Config{})
	if err != nil {
		c.Logger.Error(err)
		return "error configuring xray", err
	}
	c.Logger.Debug("X-Ray context initialized")

	result, err := businessLogicHandlerS3(ctx, c, e)

	return result, err
}
