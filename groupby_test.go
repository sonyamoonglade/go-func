package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	type Room int

	const (
		Kitchen Room = iota + 1
		Bedroom
		Lounge
	)

	type Item struct {
		room Room
		name string
	}

	var (
		bed     = Item{room: Bedroom, name: "bed"}
		table   = Item{room: Kitchen, name: "table"}
		toaster = Item{room: Kitchen, name: "toaster"}
		pillow  = Item{room: Bedroom, name: "pillow"}
		sofa    = Item{room: Lounge, name: "sofa"}
	)

	groups := GroupBy(
		[]Item{
			bed,
			table,
			toaster,
			pillow,
			sofa,
		},
		func(item Item) Room {
			return item.room
		},
	)

	assert.Equal(
		t,
		map[Room][]Item{
			Kitchen: {table, toaster},
			Bedroom: {bed, pillow},
			Lounge:  {sofa},
		},
		groups,
	)

	groups = GroupBy(nil, func(item Item) Room { return item.room })
	assert.Equal(t, map[Room][]Item{}, groups)
}
