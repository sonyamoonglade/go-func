package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastOr(t *testing.T) {
	for _, test := range firstOrAndLastOrTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.lastOr, LastOr(test.ss, 202))
		})
	}
}
