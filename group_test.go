package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	assert.Equal(t, map[float64]int{}, Group([]float64{}))

	assert.Equal(t, map[float64]int{
		1: 1,
	}, Group([]float64{1}))

	assert.Equal(t, map[float64]int{
		1: 1,
		2: 2,
		3: 3,
	}, Group([]float64{1, 2, 2, 3, 3, 3}))
}
