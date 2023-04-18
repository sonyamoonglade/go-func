package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterNot(t *testing.T) {
	for _, test := range selectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedFilterNot, FilterNot(test.ss, test.condition))
		})
	}
}
