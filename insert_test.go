package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert.Equal(t, []float64(nil), Insert([]float64(nil), 0))
	assert.Equal(t, []float64{2.0, 1.0}, Insert([]float64{1.0}, 0, 2.0))
	assert.Equal(t, []float64{1.0, 2.0}, Insert([]float64{1.0}, 1, 2.0))
	assert.Equal(t, []float64{1.0, 2.0, 3.3}, Insert([]float64{1.0, 3.3}, 1, 2.0))
}
