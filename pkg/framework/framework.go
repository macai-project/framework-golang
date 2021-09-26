package framework

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/macai-project/framework-golang/pkg/container"
	"os"
	"time"
)

var c *container.Container

var businessLogicHandler func(ctx context.Context, c *container.Container, e interface{}) (string, error)

func RegisterBusinessLogic(funzione func(ctx context.Context, c *container.Container, e interface{}) (string, error)) {
	businessLogicHandler = funzione
}

// HandleRequest avvia il framework
func HandleRequest(ctx context.Context, e interface{}) (string, error) {
	var err error

	c = &container.Container{}

	// Logger
	c.NewLogger()
	defer c.Logger.Sync()

	// AWS Config
	err = c.NewAWSConfig()
	if err != nil {
		c.Logger.Fatalf("Error initializing AWS Config: %s", err)
		return "AWS Error", err
	}

	// Sentry
	sentryDSN, _ := os.LookupEnv("SENTRY_DSN")
	err = sentry.Init(sentry.ClientOptions{
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
