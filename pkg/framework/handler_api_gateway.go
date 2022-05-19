package framework

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/getsentry/sentry-go"
	"github.com/macai-project/framework-golang/pkg/container"
	"os"
	"strconv"
	"time"
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
			StatusCode: 500,
			Body:       "AWS Config not initialized",
		}, err
	}

	// Sentry
	sentryDSN, ok := os.LookupEnv("SENTRY_DSN")
	if ok != true {
		c.Logger.Warnf("error fetching SENTRY_DSN, sentry won't sample")
		sentryDSN = ""
	}

	samplingRate, err := strconv.ParseFloat(os.Getenv("SENTRY_SAMPLING_RATE"), 64)
	if err != nil {
		c.Logger.Warnf("error converting SENTRY_SAMPLING_RATE to float64, set to 1.0")
		samplingRate = 1.0
	}

	environment, ok := os.LookupEnv("ENVIRONMENT")
	if ok != true {
		c.Logger.Warnf("error fetching ENVIRONMENT, setting empty environment")
		environment = ""
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDSN,
		Environment:      environment,
		TracesSampleRate: samplingRate,
	})

	// Xray
	awsv2.AWSV2Instrumentor(&c.AwsConfig.APIOptions)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Sentry not initialized",
		}, err
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(1 * time.Second)

	result, err := businessLogicHandlerApiGateway(ctx, c, e)
	if err != nil {
		sentry.CaptureException(err)
	}
	return result, err
}
