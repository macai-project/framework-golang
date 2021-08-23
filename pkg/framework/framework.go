package framework

import (
	"context"

	"github.com/macai-project/events"
	"github.com/macai-project/framework-golang/internal/container"
)

var businessLogicHandler func(ctx context.Context, c *container.Container, e events.TestEvent) (string, error)

func RegisterBusinessLogic(funzione func(ctx context.Context, c *container.Container, e events.TestEvent) (string, error)) {
	businessLogicHandler = funzione
}

// HandleRequest avvia il framework
func HandleRequest(ctx context.Context, c *container.Container, e events.TestEvent) (string, error) {
	return businessLogicHandler(ctx, c, e)
}
