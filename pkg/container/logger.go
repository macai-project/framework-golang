package container

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// NewLogger create a new Logger
func (c *Container) NewLogger() {

	// If the logger is already instantiated, do nothing
	if c.Logger != nil {
		return
	}

	logger, _ := zap.NewProduction(zap.Hooks(func(entry zapcore.Entry) error {
		return nil
	}))
	if os.Getenv("LOG_LEVEL") == "debug" {
		logger, _ = zap.NewDevelopment(zap.Hooks(func(entry zapcore.Entry) error {
			return nil
		}))
	}
	c.Logger = logger.Sugar()
}

// Flush buffers, if any
func (c *Container) Flush() {
	c.Logger.Sync()
}
