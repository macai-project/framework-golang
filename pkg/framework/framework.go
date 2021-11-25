package framework

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/getsentry/sentry-go"
	"github.com/macai-project/framework-golang/pkg/container"
	"os"
	"strconv"
	"time"
)

var c *container.Container

var businessLogicHandler func(ctx context.Context, c *container.Container, e events.CloudWatchEvent) (string, error)

func RegisterContainer(fc *container.Container) {
	c = fc
}

func RegisterBusinessLogic(f func(ctx context.Context, c *container.Container, e events.CloudWatchEvent) (string, error)) {
	businessLogicHandler = f
}

// HandleRequest start the framework
func HandleRequest(ctx context.Context, e events.CloudWatchEvent) (string, error) {
	var err error

	// Check if the container has been initialized
	if c == nil {
		return "", fmt.Errorf("container struct must be initialized and passed to the framework")
	}

	// Logger
	c.NewLogger()
	defer c.Logger.Sync()

	// AWS Config
	err = c.NewAWSConfig()
	if err != nil {
		return "error initializing AWS Config", err
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
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDSN,
		TracesSampleRate: samplingRate,
	})
	if err != nil {
		return "error in sentry.Init", err
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(1 * time.Second)

	result, err := businessLogicHandler(ctx, c, e)
	if err != nil {
		sentry.CaptureException(err)
	}
	return result, err
}
