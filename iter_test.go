package rbtree

import (
	"testing"

	"github.com/fanyang01/tree/common"
	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	n := 1 << 10
	tr := New(common.CompareInt)

	assert.Nil(t, tr.PostorderFirst())
	assert.Nil(t, tr.PreorderFirst())

	for i := 0; i < n; i++ {
		tr.Insert(i)
	}

	assert.Nil(t, tr.Prev(tr.First()))
	assert.Nil(t, tr.Next(tr.Last()))

	for i, x := 0, tr.First(); i < n; i++ {
		assert.NotNil(t, x)
		assert.Equal(t, i, x.Value())
		x = tr.Next(x)
	}
	for i, x := n-1, tr.Last(); i >= 0; i-- {
		assert.NotNil(t, x)
		assert.Equal(t, i, x.Value())
		x = tr.Prev(x)
	}

	tr.Clean()
	for i := 0; i < n; i++ {
		tr.Insert(r.Intn(n))
	}

	for x := tr.PostorderFirst(); x != nil; x = tr.PostorderNext(x) {
		assert.NotNil(t, tr.PostorderFirstNode(x))
	}

	for x := tr.PreorderFirst(); x != nil; x = tr.PreorderNext(x) {
		assert.NotNil(t, tr.PreorderLastNode(x))
	}
}
