package rbtree

import (
	"math/rand"
	"testing"
)

type Int int

const Count = 1 << 19

func (a Int) Compare(b Interface) int {
	y := b.(Int)
	if a > y {
		return 1
	} else if a < y {
		return -1
	}
	return 0
}

func Test(t *testing.T) {
	tr := New()
	for i := 0; i < (Count); i++ {
		I := Int(i)
		tr.Insert(I)
	}

	for i := 0; i < Count; i++ {
		I := Int(i)
		if _, err := tr.Search(I); err != nil {
			t.Fail()
		}
	}

	var deleted [Count]bool
	for i := 0; i < Count>>1; i++ {
		I := Int(rand.Intn(Count))
		if !deleted[int(I)] {
			_, err := tr.Delete(I)
			if err != nil {
				t.Fail()
			} else {
				deleted[int(I)] = true
			}
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	tr := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Insert(I)
	}
}

func BenchmarkSearch(b *testing.B) {
	tr := New()
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Insert(I)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Search(I)
	}
}
func BenchmarkDelete(b *testing.B) {
	tr := New()
	for i := 0; i < b.N; i++ {
		I := Int(i)
		tr.Insert(I)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		I := Int(rand.Intn(Count))
		tr.Delete(I)
	}
}
