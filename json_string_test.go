package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONString(t *testing.T) {
	for _, test := range jsonTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.jsonString, JSONString(test.ss))
		})
	}
}
