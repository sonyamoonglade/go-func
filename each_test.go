package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEach(t *testing.T) {
	var values []float64

	values = []float64{}
	Each([]float64{}, func(value float64) {
		values = append(values, value)
	})
	assert.Equal(t, []float64{}, values)

	values = []float64{}
	Each([]float64{435.34, 8923.1}, func(value float64) {
		values = append(values, value)
	})
	assert.Equal(t, []float64{435.34, 8923.1}, values)
}
