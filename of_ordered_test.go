package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfOrdered(t *testing.T) {
	t.Run("chaining", func(t *testing.T) {
		name := OfOrdered([]string{"Bob", "Sally", "John", "Jane"}).
			Join("+")

		assert.Equal(t, "Bob+Sally+John+Jane", name)
	})

	t.Run("result", func(t *testing.T) {
		names := OfOrdered([]string{"Bob", "Sally", "John", "Jane"}).
			Filter(func(s string) bool {
				return len(s) < 4
			}).
			Result

		assert.Equal(t, []string{"Bob"}, names)
	})
}
