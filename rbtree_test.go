package rbtree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/fanyang01/tree/common"
)

const Count = 1 << 19

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func Test(t *testing.T) {
	n := 1 << 16
	tr := New(common.CompareInt)
	if !tr.IsEmpty() {
		t.Fail()
	}
	for i := 0; i < n; i++ {
		if _, ok := tr.Insert(i); !ok {
			t.Fail()
		}
	}
	if tr.IsEmpty() {
		t.Fail()
	}
	if tr.Len() != n {
		t.Fail()
	}
	if tr.Has(-1) {
		t.Fail()
	}
	if _, ok := tr.Insert(0); ok {
		t.Fail()
	}
	for i := n - 1; i >= n/2; i-- {
		handle, _ := tr.Search(i)
		if _, ok := tr.Delete(handle); !ok {
			t.Fail()
		}
		if _, ok := tr.Search(i); ok {
			t.Fail()
		}
	}
	if tr.Len() != n-n/2 {
		t.Fail()
	}
	for i := 0; i < n/2; i++ {
		v, ok := tr.Search(i)
		if !ok || v.Value() != i {
			t.Fail()
		}
	}
	handle, _ := tr.Search(0)
	if _, ok := tr.Replace(handle, 0); !ok {
		t.Fail()
	}
	handle, _ = tr.Search(0)
	if _, ok := tr.Replace(handle, 1); ok {
		t.Fail()
	}
	deleted := make(map[int]bool)
	for i := 0; i < n/2; i++ {
		random := r.Intn(n / 2)
		var handle *Node
		var ok bool
		if handle, ok = tr.Search(random); ok {
			if _, ok := tr.Delete(handle); ok {
				if deleted[random] {
					t.Error("Already deleted")
				}
				deleted[random] = true
			}
		} else {
			if !deleted[random] {
				t.Error("Have not deleted")
			}
		}
	}
	tr.Clean()
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
		i := r.Intn(Count)
		handle, ok := tr.Search(i)
		if ok {
			tr.Delete(handle)
		}
	}
}
