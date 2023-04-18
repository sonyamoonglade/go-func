package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {
	assert.Equal(t, []int(nil), Ints([]int(nil)))

	assert.Equal(t,
		[]int{92, 823, 453},
		Ints([]float64{92.384, 823.324, 453}))
}
