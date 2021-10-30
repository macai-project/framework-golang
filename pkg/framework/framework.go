package framework

import (
	"context"
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

	c, err = container.Bootstrap()
	if err != nil {
		return err.Error(), err
	}

	// Sentry
	sentryDSN, _ := os.LookupEnv("SENTRY_DSN")
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDSN,
		TracesSampleRate: 0.2,
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
