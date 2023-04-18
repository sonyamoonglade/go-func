package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat64s(t *testing.T) {
	assert.Equal(t, []float64(nil), Float64s([]float64(nil)))

	assert.Equal(t,
		[]float64{92.384, 823.324, 453},
		Float64s([]float64{92.384, 823.324, 453}))
}
