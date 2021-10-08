package container

import (
	"go.uber.org/zap"
	"os"
)

// NewLogger create a new Logger
func (c *Container) NewLogger() {
	logger, _ := zap.NewProduction()
	if os.Getenv("LOG_LEVEL") == "debug" {
		logger, _ = zap.NewDevelopment()
	}
	c.Logger =logger.Sugar()
}

// Flush buffers, if any
func (c *Container) Flush() {
	c.Logger.Sync()
}
