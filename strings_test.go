package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
	assert.Equal(t, []string{}, Strings([]float64{}))

	assert.Equal(t,
		[]string{"92.384", "823.324", "453"},
		Strings([]float64{92.384, 823.324, 453}))
}
