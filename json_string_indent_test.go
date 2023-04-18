package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONStringIndent(t *testing.T) {
	for _, test := range jsonIndentTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.jsonString, JSONStringIndent(test.ss, "", "  "))
		})
	}
}
