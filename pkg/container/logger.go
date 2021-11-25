package container

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// NewLogger create a new Logger
func (c *Container) NewLogger() {

	// If the logger is already instantiated, do nothing
	if c.Logger != nil {
		return
	}

	logger, _ := zap.NewProduction(zap.Hooks(func(entry zapcore.Entry) error {
		if entry.Level == zapcore.WarnLevel {
			defer sentry.Flush(2 * time.Second)
			sentry.CaptureMessage(fmt.Sprintf("%s, Line No: %d :: %s", entry.Caller.File, entry.Caller.Line, entry.Message))
		}
		return nil
	}))
	if os.Getenv("LOG_LEVEL") == "debug" {
		logger, _ = zap.NewDevelopment(zap.Hooks(func(entry zapcore.Entry) error {
			if entry.Level == zapcore.WarnLevel {
				defer sentry.Flush(2 * time.Second)
				sentry.CaptureMessage(fmt.Sprintf("%s, Line No: %d :: %s", entry.Caller.File, entry.Caller.Line, entry.Message))
			}
			return nil
		}))
	}
	c.Logger = logger.Sugar()
}

// Flush buffers, if any
func (c *Container) Flush() {
	c.Logger.Sync()
}
