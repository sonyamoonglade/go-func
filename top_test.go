package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTop(t *testing.T) {
	for _, test := range topAndBottomTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.top, Top(test.ss, test.n))
		})
	}
}
