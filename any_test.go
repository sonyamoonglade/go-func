package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	assert.False(t,
		Any([]float64{}, func(value float64) bool {
			return true
		}),
	)

	// None
	assert.False(t,
		Any([]float64{12.3, 4.56}, func(value float64) bool {
			return value == 1
		}),
	)

	// Some
	assert.True(t,
		Any([]float64{12.3, 4.56}, func(value float64) bool {
			return value == 12.3
		}),
	)

	// All
	assert.True(t,
		Any([]float64{12.3, 4.56}, func(value float64) bool {
			return value > 0
		}),
	)
}
