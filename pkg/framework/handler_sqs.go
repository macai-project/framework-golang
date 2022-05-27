package framework

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/macai-project/framework-golang/pkg/container"
)

var businessLogicHandlerSQS func(ctx context.Context, c *container.Container, e events.SQSEvent) (string, error)

func RegisterBusinessLogicSQS(f func(ctx context.Context, c *container.Container, e events.SQSEvent) (string, error)) {
	businessLogicHandlerSQS = f
}

// HandleRequestSQS start the framework
func HandleRequestSQS(ctx context.Context, e events.SQSEvent) (string, error) {
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

	result, err := businessLogicHandlerSQS(ctx, c, e)
	return result, err
}
