package go_func

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	assert.Equal(t, []currency(nil), Values(currencies(nil)))

	assert.Equal(t, []currency(nil), Values(currencies{}))

	values := Values(isoCurrencies)
	sort.Slice(values, func(i, j int) bool {
		return values[i].NumericCode < values[j].NumericCode
	})

	assert.Equal(t, []currency{{36, -2}, {840, -2}}, values)
}
