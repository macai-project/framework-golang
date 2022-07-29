package framework

import (
	"context"
	"fmt"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/macai-project/framework-golang/pkg/container"
)

var businessGenericLogicHandler func(ctx context.Context, c *container.Container, e interface{}) (string, error)

func RegisterGenericBusinessLogic(f func(ctx context.Context, c *container.Container, e interface{}) (string, error)) {
	businessGenericLogicHandler = f
}

func HandleGenericRequest(ctx context.Context, e interface{}) (string, error) {
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

	result, err := businessGenericLogicHandler(ctx, c, e)

	return result, err
}
