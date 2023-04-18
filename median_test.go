package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedian(t *testing.T) {
	assert.Equal(t, 0.0, Median([]float64{}))
	assert.Equal(t, 12.3, Median([]float64{12.3}))
	assert.Equal(t, 8.4, Median([]float64{12.3, 4.5}))
	assert.Equal(t, 4.5, Median([]float64{2.1, 12.3, 4.5}))
}
