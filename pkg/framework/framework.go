package framework

import (
	"context"
)

var businessLogicHandler func(ctx context.Context, e testEvent) (string, error)

func RegisterBusinessLogic(funzione func(ctx context.Context, e testEvent) (string, error)) {
	businessLogicHandler = funzione
}

// HandleRequest avvia il framework
func HandleRequest(ctx context.Context, e testEvent) (string, error) {
	return businessLogicHandler(ctx, e)
}
