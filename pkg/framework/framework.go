package framework

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

var businessLogicHandler func(ctx context.Context, e events.CloudWatchEvent) (string, error)

func RegisterBusinessLogic(funzione func(ctx context.Context, e events.CloudWatchEvent) (string, error)) {
	businessLogicHandler = funzione
}

// HandleRequest avvia il framework
func HandleRequest(ctx context.Context, e events.CloudWatchEvent) (string, error) {
	return businessLogicHandler(ctx, e)
}
