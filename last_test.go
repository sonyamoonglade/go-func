package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLast(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.last, Last(test.ss))
		})
	}
}
