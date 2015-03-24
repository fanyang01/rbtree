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
	rbt := New()
	for i := 0; i < (Count); i++ {
		I := Int(i)
		rbt.Insert(I)
	}
	for i := 0; i < Count; i++ {
		I := Int(i)
		if _, err := rbt.Search(I); err != nil {
			t.Fail()
		}
	}
	var deleted [Count]bool
	for i := 0; i < Count>>1; i++ {
		I := Int(rand.Intn(Count))
		if !deleted[int(I)] {
			_, err := rbt.Delete(I)
			if err != nil {
				log.Println(int(I), err.Error())
				t.Fail()
			} else {
				deleted[int(I)] = true
			}
		}
	}
}
