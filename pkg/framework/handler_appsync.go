package framework

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-xray-sdk-go/instrumentation/awsv2"
	"github.com/getsentry/sentry-go"
	"github.com/macai-project/framework-golang/pkg/container"
)

var businessLogicHandlerAppsync func(ctx context.Context, c *container.Container, e events.AppSyncResolverTemplate) (interface{}, error)

func RegisterBusinessLogicAppsync(f func(ctx context.Context, c *container.Container, e events.AppSyncResolverTemplate) (interface{}, error)) {
	businessLogicHandlerAppsync = f
}

// HandleRequestAppsync start the framework
func HandleRequestAppsync(ctx context.Context, e events.AppSyncResolverTemplate) (interface{}, error) {
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
	if err != nil {
		return "Sentry not initialized", err
	}

	// Xray
	awsv2.AWSV2Instrumentor(&c.AwsConfig.APIOptions)

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(1 * time.Second)

	result, err := businessLogicHandlerAppsync(ctx, c, e)
	if err != nil {
		sentry.CaptureException(err)
	}
  
	return result, err
}
