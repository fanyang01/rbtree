package rbtree

import (
	"testing"

	"github.com/fanyang01/tree/common"
	"github.com/stretchr/testify/assert"
)

func TestWalk(t *testing.T) {
	n := 1 << 10
	tr := New(common.CompareInt)
	for i := 0; i < n; i++ {
		tr.Insert(i)
	}

	var i int
	fn := func(x *Node) bool {
		assert.Equal(t, i, x.Value())
		i++
		return true
	}
	tr.Walk(VisitFunc(fn))

	i = n - 1
	fn = func(x *Node) bool {
		assert.Equal(t, i, x.Value())
		i--
		return true
	}
	tr.WalkReverse(VisitFunc(fn))

	fn = func(x *Node) bool {
		left, right := x.Left(), x.Right()
		if left != nil {
			assert.True(t, x.Value().(int) > left.Value().(int))
		}
		if right != nil {
			assert.True(t, x.Value().(int) < right.Value().(int))
		}
		return true
	}
	tr.WalkPostorder(VisitFunc(fn))
	tr.WalkPreorder(VisitFunc(fn))

	size := 0
	fn = func(x *Node) bool {
		size++
		return true
	}
	tr.WalkPostorder(VisitFunc(fn))
	assert.Equal(t, n, size)

	size = 0
	tr.WalkSubPostorder(VisitFunc(fn), tr.Root())
	assert.Equal(t, n, size)

	size = 0
	tr.WalkPreorder(VisitFunc(fn))
	assert.Equal(t, n, size)

	size = 0
	tr.WalkSubPreorder(VisitFunc(fn), tr.Root())
	assert.Equal(t, n, size)

	size = 0
	fn = func(x *Node) bool {
		size++
		return false
	}
	tr.Walk(VisitFunc(fn))
	tr.WalkReverse(VisitFunc(fn))
	tr.WalkPostorder(VisitFunc(fn))
	tr.WalkPreorder(VisitFunc(fn))
	tr.WalkSubPostorder(VisitFunc(fn), tr.Root())
	tr.WalkSubPreorder(VisitFunc(fn), tr.Root())
	assert.Equal(t, 6, size)
}
