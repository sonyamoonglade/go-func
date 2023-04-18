package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceUsing(t *testing.T) {
	for _, test := range sequenceAndSequenceUsingTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, SequenceUsing(test.ss,
				func(i int) float64 { return float64(i) }, test.params...))
		})
	}
}
