package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	assert.Equal(t, 123, Int(123))
	assert.Equal(t, 1, Int(1.89))
	assert.Equal(t, 1, Int("1.89"))
}
