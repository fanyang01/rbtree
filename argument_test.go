package rbtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// maximum value
	argFunc = func(n *Node) {
		right := n.Right()
		if right == nil {
			n.Argument = n.Value()
			return
		}
		n.Argument = right.Argument
	}

	cmpArg = func(x, y interface{}) int {
		if x == y {
			return 0
		}
		return 1
	}
)

func TestArg(t *testing.T) {
	n := 1 << 10
	tr := New(compareInt, cmpArg, argFunc)
	for i := 0; i < n; i++ {
		tr.Insert(i)
		assert.Equal(t, i, tr.Root().Argument)
	}

	for i := n - 1; i > 0; i-- {
		tr.DeleteValue(i)
		assert.Equal(t, i-1, tr.Root().Argument)
	}
}
