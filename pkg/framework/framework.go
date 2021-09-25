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

func RegisterBusinessLogic(funzione func(ctx context.Context, c *container.Container, e events.CloudWatchEvent) (string, error)) {
	businessLogicHandler = funzione
}

// HandleRequest avvia il framework
func HandleRequest(ctx context.Context, e events.CloudWatchEvent) (string, error) {
	c = &container.Container{}

	// Sentry
	sentryDSN, _ := os.LookupEnv("SENTRY_DSN")
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDSN,
		TracesSampleRate: 0.2,
	})
	if err != nil {
		c.Logger.Fatalf("sentry.Init: %s", err)
		return "sentry.Init", err
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(1 * time.Second)

	return businessLogicHandler(ctx, c, e)
}
