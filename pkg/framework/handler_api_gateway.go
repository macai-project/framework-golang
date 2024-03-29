package framework

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/macai-project/framework-golang/pkg/container"
	"net/http"
)

var businessLogicHandlerApiGateway func(ctx context.Context, c *container.Container, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func RegisterBusinessLogicApiGateway(f func(ctx context.Context, c *container.Container, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) {
	businessLogicHandlerApiGateway = f
}

// HandleRequestApiGateway start the framework
func HandleRequestApiGateway(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var err error

	// Check if the container has been initialized
	if c == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Container not initialized",
		}, fmt.Errorf("container struct must be initialized and passed to the framework")
	}

	// Logger
	c.NewLogger()
	defer c.Logger.Sync()

	// AWS Config
	err = c.NewAWSConfig(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "AWS Config not initialized",
		}, err
	}

	// Xray
	awsv2.AWSV2Instrumentor(&c.AwsConfig.APIOptions)

	ctx, err = xray.ContextWithConfig(ctx, xray.Config{})
	if err != nil {
		c.Logger.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "error configuring xray",
		}, err
	}
	c.Logger.Debug("X-Ray context initialized")

	result, err := businessLogicHandlerApiGateway(ctx, c, e)

	return result, err
}
