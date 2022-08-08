package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMidecFromEAN(t *testing.T) {
	midec := MidecFromEAN("3259920039721")
	assert.Equal(t, midec, "12bf955e-6b77-5831-87d3-9ead699e0aeb")
}
