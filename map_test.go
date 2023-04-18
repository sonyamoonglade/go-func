package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	for _, test := range selectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedMap, Map(test.ss, func(a float64) float64 {
				return a + 5.2
			}))
		})
	}
}
