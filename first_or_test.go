package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var firstOrAndLastOrTests = []struct {
	ss      []float64
	firstOr float64
	lastOr  float64
}{
	{
		nil,
		102,
		202,
	},
	{
		[]float64{100},
		100,
		100,
	},
	{
		[]float64{1, 2},
		1,
		2,
	},
	{
		[]float64{1, 2, 3},
		1,
		3,
	},
}

func TestFirstOr(t *testing.T) {
	for _, test := range firstOrAndLastOrTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.firstOr, FirstOr(test.ss, 102))
		})
	}
}
