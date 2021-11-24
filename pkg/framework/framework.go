package framework

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/getsentry/sentry-go"
	"github.com/macai-project/framework-golang/pkg/container"
	"os"
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

	// AWS Config V1
	err = c.NewAWSConfigV1()
	if err != nil {
		return "error initializing AWS Config V1", err
	}

	// AWS Config V2
	err = c.NewAWSConfig()
	if err != nil {
		return "error initializing AWS Config", err
	}

	// Sentry
	sentryDSN, _ := os.LookupEnv("SENTRY_DSN")
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDSN,
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
