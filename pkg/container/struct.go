package container

import (
	"go.uber.org/zap"
)

// Container is the container used to inject dependencies into the flow
type Container struct {
	Logger *zap.SugaredLogger
}
