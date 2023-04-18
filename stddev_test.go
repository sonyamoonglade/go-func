package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStddev(t *testing.T) {
	assert.Equal(t, 0.0, Stddev([]float64{}))
	assert.Equal(t, 0.0, Stddev([]float64{1}))
	assert.Equal(t, 4.8587389053127765, Stddev([]float64{10.0, 12.5, 23.3, 23.1, 16.5, 23.1, 21.2, 16.4}))
}
