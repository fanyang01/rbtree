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
	n := 1 << 10
	tr := New(common.CompareInt)
	for i := 0; i < n; i++ {
		if !tr.Insert(i) {
			t.Fail()
		}
	}
	if !checkRbTree(tr) {
		t.Error("Not a valid red black tree")
	}
	for i := 0; i < n; i++ {
		v := rand.Intn(n)
		tr.Delete(v)
	}
	if !checkRbTree(tr) {
		t.Error("Not a valid red black tree")
	}
}

func checkRbTree(t *RbTree) bool {
	if t.root.color != BLACK || t.null.color != BLACK {
		return false
	}
	h := new(Height)
	return checkColor(t, t.root) && h.checkHeight(t, t.root, 0)
}

func checkColor(t *RbTree, n *node) bool {
	checkNode := func(n *node) bool {
		if n.color == RED {
			if n.left.color != BLACK || n.right.color != BLACK {
				return false
			}
		}
		return true
	}
	ok := checkNode(n)
	if !ok || n == t.null {
		return ok
	}
	return checkColor(t, n.left) && checkColor(t, n.right)
}

func (h *Height) checkHeight(t *RbTree, n *node, height int) bool {
	if n == t.null {
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
