/*
Package rbtree implements red-black tree introduced in "Introduction to Algorithms".
*/
package rbtree

import "github.com/fanyang01/tree/common"

// BLACK and RED is the color of nodes
const (
	BLACK = false
	RED   = true
)

// Node is the node in a tree
type Node struct {
	left, right, p *Node
	color          bool
	v              interface{}
}

// Tree is a red-black tree
type Tree struct {
	size    int
	root    *Node
	compare common.CompareFunc
}

// Left returns the left child of n
func (n *Node) Left() *Node { return n.left }

// Right returns the right child of n
func (n *Node) Right() *Node { return n.right }

// Parent returns the parent of n
func (n *Node) Parent() *Node { return n.p }

// Value returns payload contained in n
func (n *Node) Value() interface{} { return n.v }

// New creates an initialized tree.
func New(f common.CompareFunc) *Tree {
	return &Tree{
		size:    0,
		root:    nil,
		compare: f,
	}
}

// IsEmpty returns true if the tree is empty.
func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

// Len returns size of t.
func (t *Tree) Len() int {
	return t.size
}

// Clean resets a tree structure to it's initial state.
func (t *Tree) Clean() *Tree {
	t.size = 0
	t.root = nil
	return t
}

// Has tests if v is already in t.
func (t *Tree) Has(v interface{}) bool {
	return t.search(t.root, v) != nil
}

// Replace replaces payload of a node with v.
// v must be equal to previous payload.
func (t *Tree) Replace(n *Node, v interface{}) (interface{}, bool) {
	if t.compare(n.v, v) != 0 {
		return n.v, false
	}
	before := n.v
	n.v = v
	return before, true
}

// Search tries to find the node containing payload v.
// On success, the node containing v will be returned,
// otherwise, nil will be returned to indicate the node is not found.
func (t *Tree) Search(v interface{}) *Node {
	return t.search(t.root, v)
}

func (t *Tree) search(r *Node, v interface{}) *Node {
	x := r
	for x != nil {
		var cmp int
		if cmp = t.compare(v, x.v); cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x
		}
	}
	return nil
}

// Insert inserts v into correct place and returns a handle.
// It will refuse to insert v when v is already in t, and returns the node.
func (t *Tree) Insert(v interface{}) (*Node, bool) {
	var cmp int
	var p *Node
	x := t.root

	for x != nil {
		p = x
		if cmp = t.compare(v, x.v); cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			// Disable duplicate v
			return x, false
		}
	}

	n := t.newNode(v)
	n.p = p
	if p == nil {
		t.root = n
	} else if cmp = t.compare(v, p.v); cmp < 0 {
		p.left = n
	} else {
		p.right = n
	}
	t.insertFix(n)
	t.size++
	return n, true
}

// DeleteValue deletes the node whose payload is equal to v.
// A boolean value is returned to indicate whether the node is found.
func (t *Tree) DeleteValue(v interface{}) (interface{}, bool) {
	if x := t.Search(v); x != nil {
		return t.Delete(x), true
	}
	return nil, false
}

// Delete removes x from t and returns its payload.
func (t *Tree) Delete(x *Node) interface{} {
	// z is the node that is MOVED to a new place,
	// and color is the color of the node previously in this place.
	var z, p *Node
	color := x.color

	if x.left == nil {
		z, p = x.right, x.p
		t.transplant(x, x.right)
	} else if x.right == nil {
		z, p = x.left, x.p
		t.transplant(x, x.left)
	} else {
		// y is the maximum node on x's right subtree
		// it will replace x
		y := func(n *Node) *Node {
			for n.left != nil {
				n = n.left
			}
			return n
		}(x.right)

		color = y.color
		// NOTE: it's important to update p to point to parent of y.right
		z = y.right
		// Avoid y.p to point to y itself
		if x.right == y {
			p = y
		} else {
			t.transplant(y, y.right)
			p = y.p
			y.right = x.right
			x.right.p = y
		}
		y.left = x.left
		x.left.p = y
		t.transplant(x, y)
		y.color = x.color
	}
	if color == BLACK {
		t.deleteFix(p, z)
	}
	t.size--
	return x.v
}
