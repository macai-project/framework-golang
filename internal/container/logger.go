package container

import "go.uber.org/zap"

// NewLogger create a new Logger
func (c *Container) NewLogger() {
	logger, _ := zap.NewProduction()
	c.Logger = logger.Sugar()
}

// Flush buffers, if any
func (c *Container) Flush() {
	c.Logger.Sync()
}
