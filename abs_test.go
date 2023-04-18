package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 584.2727, Abs(-584.2727))
	assert.Equal(t, 5, Abs(-5))
	assert.Equal(t, 584.2727, Abs(584.2727))
}
