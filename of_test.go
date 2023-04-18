package go_func

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {

	t.Run("result", func(t *testing.T) {
		names := Of([]string{"Bob", "Sally", "John", "Jane"}).
			FilterNot(func(name string) bool {
				return strings.HasPrefix(name, "J")
			}).
			Result

		assert.Equal(t, []string{"Bob", "Sally"}, names)
	})
}
