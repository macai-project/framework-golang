package framework

import (
	"context"
	"fmt"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/macai-project/framework-golang/pkg/container"
)

var businessLogicHandlerAppsync func(ctx context.Context, c *container.Container, e map[string]interface{}) (interface{}, error)

func RegisterBusinessLogicAppsync(f func(ctx context.Context, c *container.Container, e map[string]interface{}) (interface{}, error)) {
	businessLogicHandlerAppsync = f
}

// HandleRequestAppsync start the framework
func HandleRequestAppsync(ctx context.Context, e map[string]interface{}) (interface{}, error) {
	var err error

	// Check if the container has been initialized
	if c == nil {
		return "Container not initialized", fmt.Errorf("container struct must be initialized and passed to the framework")
	}

	// Logger
	c.NewLogger()
	defer c.Logger.Sync()

	// AWS Config
	err = c.NewAWSConfig(ctx)
	if err != nil {
		return "AWS Config not initialized", err
	}

	// Xray
	awsv2.AWSV2Instrumentor(&c.AwsConfig.APIOptions)

	ctx, err = xray.ContextWithConfig(ctx, xray.Config{})
	if err != nil {
		c.Logger.Error(err)
		return "error configuring xray", err
	}
	c.Logger.Debug("X-Ray context initialized")

	result, err := businessLogicHandlerAppsync(ctx, c, e)

	return result, err
}
