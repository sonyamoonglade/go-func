package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfONumeric(t *testing.T) {
	t.Run("chaining", func(t *testing.T) {
		total := OfNumeric([]float64{123, 456}).
			Sum()

		assert.Equal(t, 579.0, total)
	})

	t.Run("result", func(t *testing.T) {
		names := OfNumeric([]float64{1.23, 4.56}).
			Filter(func(x float64) bool {
				return x < 4
			}).
			Result

		assert.Equal(t, []float64{1.23}, names)
	})
}
