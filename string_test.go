package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "123", String(123))
	assert.Equal(t, "1.89", String(1.89))
	assert.Equal(t, "1.89", String("1.89"))
}
