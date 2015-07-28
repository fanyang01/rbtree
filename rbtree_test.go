package rbtree

import (
	"math/rand"
	"testing"

	"github.com/fanyang01/tree/common"
)

const Count = 1 << 19

func Test(t *testing.T) {
	n := 1 << 16
	tr := New(common.CompareInt)
	if !tr.IsEmpty() {
		t.Fail()
	}
	for i := 0; i < n; i++ {
		if !tr.Insert(i) {
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
	if tr.Insert(0) {
		t.Fail()
	}
	if _, ok := tr.Delete(n + 1); ok {
		t.Fail()
	}
	for i := n - 1; i > n/2; i-- {
		if _, ok := tr.Delete(i); !ok {
			t.Fail()
		}
		if _, ok := tr.Search(i); ok {
			t.Fail()
		}
	}
	for i := 0; i <= n/2; i++ {
		v, ok := tr.Search(i)
		if !ok || v != i {
			t.Fail()
		}
	}
	if _, ok := tr.Replace(0); !ok {
		t.Fail()
	}
	if _, ok := tr.Replace(-1); ok {
		t.Fail()
	}
	deleted := make(map[int]bool)
	for i := 0; i <= n/2; i++ {
		random := rand.Intn(n/2 + 1)
		if _, ok := tr.Delete(random); ok {
			if deleted[random] {
				t.Error("Already deleted")
			}
			deleted[random] = true
		} else {
			if !deleted[random] {
				t.Error("Have not deleted")
			}
		}
	}
	tr.Clean()
	for i := 0; i < n; i++ {
		random := rand.Intn(n)
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
		i := rand.Intn(Count)
		tr.Delete(i)
	}
}
