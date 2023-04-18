package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	assert.True(t,
		All([]float64{}, func(value float64) bool {
			return false
		}),
	)

	// None
	assert.False(t,
		All([]float64{12.3, 4.56}, func(value float64) bool {
			return value == 1
		}),
	)

	// Some
	assert.False(t,
		All([]float64{12.3, 4.56}, func(value float64) bool {
			return value == 12.3
		}),
	)

	// All
	assert.True(t,
		All([]float64{12.3, 4.56}, func(value float64) bool {
			return value > 0
		}),
	)
}
