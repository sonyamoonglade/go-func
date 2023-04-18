package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	for _, test := range sortTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.reversed, Reverse(test.ss))
		})
	}
}
