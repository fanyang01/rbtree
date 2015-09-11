package rbtree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func TestTree(t *testing.T) {
	n := 1 << 16
	tr := New(compareInt, cmpArg, argFunc)

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

	tr.Insert(1)
	assert.Nil(t, tr.Root().Parent())
	assert.Nil(t, tr.Root().Left())
	assert.Nil(t, tr.Root().Right())
	tr.Clean()

	for i := 0; i < n; i++ {
		random := r.Intn(n)
		tr.Insert(random)
	}
}

func BenchmarkInsert(b *testing.B) {
	tr := New(compareInt, cmpArg, argFunc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Insert(i)
	}
}

func BenchmarkSearch(b *testing.B) {
	tr := New(compareInt, cmpArg, argFunc)
	for i := 0; i < b.N; i++ {
		tr.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Search(i)
	}
}

func BenchmarkDelete(b *testing.B) {
	tr := New(compareInt, cmpArg, argFunc)
	for i := 0; i < b.N; i++ {
		tr.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := r.Intn(b.N)
		tr.DeleteValue(v)
	}
}
