package rbtree

import (
	"log"
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

func TestNew(t *testing.T) {
	treap := New()
	for i := 0; i < (Count); i++ {
		I := Int(i)
		treap.Insert(I)
	}
	var deleted [Count]bool
	for i := 0; i < Count; i++ {
		I := Int(rand.Intn(Count))
		if !deleted[int(I)] {
			_, err := treap.Delete(I)
			if err != nil {
				log.Println(int(I), err.Error())
				t.Fail()
			} else {
				deleted[int(I)] = true
			}
		}
	}
}
