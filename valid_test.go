package rbtree

import (
	"math/rand"
	"testing"

	"github.com/fanyang01/tree/common"
)

type Height struct {
	set    bool
	height int
}

func TestValid(t *testing.T) {
	n := 1 << 16
	tr := New(common.CompareInt)
	for i := 0; i < n; i++ {
		tr.Insert(i)
	}
	if !checkRbTree(tr) {
		t.Error("Not a valid red black tree")
	}
	for i := 0; i < n; i++ {
		v := rand.Intn(n)
		tr.DeleteValue(v)
	}
	if !checkRbTree(tr) {
		t.Error("Not a valid red black tree")
	}
}

func checkRbTree(t *Tree) bool {
	if t.root != nil && t.root.color != BLACK {
		return false
	}
	h := new(Height)
	// for performance reason, only check black-height of root
	return checkColor(t, t.root) && h.checkHeight(t, t.root, 0)
}

func checkColor(t *Tree, n *Node) bool {
	if n == nil {
		return true
	}
	checkNode := func(x *Node) bool {
		if x.color == RED {
			if !isBlack(x.left) || !isBlack(x.right) {
				return false
			}
		}
		return true
	}
	ok := checkNode(n)
	if !ok {
		return false
	}
	return checkColor(t, n.left) && checkColor(t, n.right)
}

func (h *Height) checkHeight(t *Tree, n *Node, height int) bool {
	if n == nil {
		if h.set {
			if height != h.height {
				return false
			}
		} else {
			h.height = height
			h.set = true
		}
		return true
	}
	if n.color == BLACK {
		height++
	}
	return h.checkHeight(t, n.left, height) && h.checkHeight(t, n.right, height)
}
