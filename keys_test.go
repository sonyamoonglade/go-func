package go_func

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type currency struct {
	NumericCode, Exponent int
}

type currencies map[string]currency

var isoCurrencies = currencies{
	"AUD": {36, -2},
	"USD": {840, -2},
}

func TestKeys(t *testing.T) {
	assert.Equal(t, []string(nil), Keys(currencies(nil)))

	assert.Equal(t, []string(nil), Keys(currencies{}))

	keys := Keys(isoCurrencies)
	sort.Strings(keys)
	assert.Equal(t, []string{"AUD", "USD"}, keys)
}
