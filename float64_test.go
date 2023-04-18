package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat64(t *testing.T) {
	assert.Equal(t, 123.0, Float64(123))
	assert.Equal(t, 1.89, Float64(1.89))
	assert.Equal(t, 1.89, Float64("1.89"))
}
