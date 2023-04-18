package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	for _, test := range uniqueTests {
		t.Run("", func(t *testing.T) {
			// We have to sort the unique slice because it is always returned in
			// random order.
			assert.Equal(t, test.unique, Sort(Unique(test.ss)))
		})
	}
}
