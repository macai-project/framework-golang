package framework

import (
	"context"
)

var businessLogicHandler func(ctx context.Context, e SupermercacchioEvent) (string, error)

func RegisterBusinessLogic(funzione func(ctx context.Context, e SupermercacchioEvent) (string, error)) {
	businessLogicHandler = funzione
}

// HandleRequest avvia il framework
func HandleRequest(ctx context.Context, e SupermercacchioEvent) (string, error) {
	return businessLogicHandler(ctx, e)
}
