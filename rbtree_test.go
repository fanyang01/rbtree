package rbtree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fanyang01/tree/common"
	"github.com/stretchr/testify/assert"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestAssert(t *testing.T) {
	n := 1 << 16
	tr := New(common.CompareInt)
	assert.True(t, tr.IsEmpty())

	for i := 0; i < n; i++ {
		_, ok := tr.Insert(i)
		assert.True(t, ok)
	}
	assert.False(t, tr.IsEmpty())
	assert.Equal(t, n, tr.Len())
	assert.False(t, tr.Has(-1))

	_, ok := tr.Insert(0)
	assert.False(t, ok)

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

	for i := n - 1; i >= n/2; i-- {
		tr.Delete(tr.Search(i))
		assert.Nil(t, tr.Search(i))
	}
	assert.Equal(t, n-n/2, tr.Len())

	for i := 0; i < n/2; i++ {
		v := tr.Search(i)
		assert.NotNil(t, v)
		assert.Equal(t, i, v.Value())
	}

	_, ok = tr.Replace(tr.Search(0), 0)
	assert.True(t, ok)
	_, ok = tr.Replace(tr.Search(0), 1)
	assert.False(t, ok)

	deleted := make(map[int]bool)
	for i := 0; i < n/2; i++ {
		random := r.Intn(n / 2)
		if _, ok := tr.DeleteValue(random); ok {
			assert.False(t, deleted[random])
			deleted[random] = true
		} else {
			assert.True(t, deleted[random])
		}
	}

	tr.Clean()
	assert.Equal(t, 0, tr.Len())
	assert.Nil(t, tr.First())
	assert.Nil(t, tr.Last())

	for i := 0; i < n; i++ {
		random := r.Intn(n)
		tr.Insert(random)
	}
}

func BenchmarkInsert(b *testing.B) {
	tr := New(common.CompareInt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Insert(i)
	}
}

func BenchmarkSearch(b *testing.B) {
	tr := New(common.CompareInt)
	for i := 0; i < b.N; i++ {
		tr.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Search(i)
	}
}

func BenchmarkDelete(b *testing.B) {
	tr := New(common.CompareInt)
	for i := 0; i < b.N; i++ {
		tr.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := r.Intn(b.N)
		tr.DeleteValue(v)
	}
}
