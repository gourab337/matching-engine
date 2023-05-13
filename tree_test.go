package orderbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {

	tree := limitPriceTree{}
	tree.addLimit(1000, &order{001, true, 1000, 5, &order{}, &order{}})
	tt := tree.addLimit(1010, &order{002, true, 1000, 5, &order{}, &order{}})
	tt.orders.add(&order{003, true, 1000, 5, &order{}, &order{}})

	tree.addLimit(1010, &order{004, true, 1000, 5, &order{}, &order{}})
	tree.addLimit(9010, &order{005, true, 1000, 5, &order{}, &order{}})
	tree.addLimit(900, &order{006, true, 1000, 5, &order{}, &order{}})
	tree.addLimit(1009, &order{007, true, 1000, 5, &order{}, &order{}})
	tree.addLimit(1008, &order{007, true, 1000, 5, &order{}, &order{}})

	s := tree.string()
	t.Log(s)

	cool := true

	assert.True(t, cool)
}
