package container

import (
	"fmt"
)

// Bootstrap setup container defaults
func Bootstrap() (*Container, error) {
	c := &Container{}

	// Logger
	c.NewLogger()
	defer c.Logger.Sync()

	// AWS Config
	err := c.NewAWSConfig()
	if err != nil {
		return nil, fmt.Errorf("error initializing AWS Config (%w", err)
	}

	return c, nil
}
