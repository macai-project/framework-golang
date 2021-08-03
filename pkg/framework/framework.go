package framework

import (
	"context"

	"github.com/macai-project/events"
)

var businessLogicHandler func(ctx context.Context, e events.TestEvent) (string, error)

func RegisterBusinessLogic(funzione func(ctx context.Context, e events.TestEvent) (string, error)) {
	businessLogicHandler = funzione
}

// HandleRequest avvia il framework
func HandleRequest(ctx context.Context, e events.TestEvent) (string, error) {
	return businessLogicHandler(ctx, e)
}
