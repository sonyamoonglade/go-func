package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	numbers := []float64{42.0, 4.2}

	assert.Equal(t, 42.0, *Pop(&numbers))
	assert.Equal(t, []float64{4.2}, numbers)

	assert.Equal(t, 4.2, *Pop(&numbers))
	assert.Equal(t, []float64{}, numbers)
}
