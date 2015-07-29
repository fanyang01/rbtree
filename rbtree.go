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

// Value returns payload contained in n
func (n *Node) Value() interface{} {
	return n.v
}

// Tree is a red-black tree
type Tree struct {
	size       int
	null, root *Node
	compare    common.CompareFunc
}

// New creates an initialized tree.
func New(f common.CompareFunc) *Tree {
	null := new(Node)
	null.p = null
	null.right = null
	null.left = null
	null.color = BLACK
	return &Tree{
		size:    0,
		null:    null,
		root:    null,
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
	t.root = t.null
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
	for x != t.null {
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
	x := t.root
	n := t.newNode(v)
	p := t.null
	var cmp int
	for x != t.null {
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
	n.p = p
	if p == t.null {
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

/*
 * 1.                B
 *                  / \
 *                 R   R  <- y
 *                /
 *         x ->  R
 * 2.
 *                   B
 *                 /   \
 *                R     B  <- y
 *                 \
 *            x ->  R
 * 3.
 *                   B
 *                 /   \
 *                R     B  <- y
 *               /
 *        x ->  R
 */
func (t *Tree) insertFix(x *Node) {
	var y *Node

	for x.p.color == RED {
		if x.p == x.p.p.left {
			y = x.p.p.right
			if y.color == RED {
				x.p.color = BLACK
				y.color = BLACK
				x.p.p.color = RED
				x = x.p.p
			} else {
				if x == x.p.right {
					x = x.p
					t.leftRotate(x)
				}
				x.p.color = BLACK
				x.p.p.color = RED
				t.rightRotate(x.p.p)
			}
		} else {
			y = x.p.p.left
			if y.color == RED {
				x.p.color = BLACK
				y.color = BLACK
				x.p.p.color = RED
				x = x.p.p
			} else {
				if x == x.p.left {
					x = x.p
					t.rightRotate(x)
				}
				x.p.color = BLACK
				x.p.p.color = RED
				t.leftRotate(x.p.p)
			}
		}
	}
	t.root.color = BLACK
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
	var z *Node
	y := x
	color := x.color
	if x.left == t.null {
		z = x.right
		t.transplant(x, x.right)
	} else if x.right == t.null {
		z = x.left
		t.transplant(x, x.left)
	} else {
		// y is the maximum node on x's right subtree
		// it will replace x
		y = func(n *Node) *Node {
			for n.left != t.null {
				n = n.left
			}
			return n
		}(x.right)

		color = y.color
		// z will replace y
		z = y.right
		if x.right == y {
			// Avoid y.p to point to y itself
			// Following statment may seem unneccesary,
			// but when z is t.null, value of it's p is uncertain
			z.p = y
		} else {
			t.transplant(y, y.right)
			y.right = x.right
			x.right.p = y
		}
		y.left = x.left
		x.left.p = y
		t.transplant(x, y)
		y.color = x.color
	}
	if color == BLACK {
		t.deleteFix(z)
	}
	t.size--
	return x.v
}

/*
 * when this function is called, x seems to have an additional
 * black color, so this function aims to shift the extra color.
 * 1.
 *                    B
 *                 /     \
 *          x ->  B       R  <- y
 *                       / \
 *                      B   B
 * 2.
 *                   R|B
 *                 /     \
 *          x ->  B       B  <- y
 *                      /   \
 *                     B     B
 * 3.
 *                   R|B
 *                 /     \
 *          x ->  B       B  <- y
 *                      /   \
 *                     R     B
 * 4.
 *                   R|B
 *                 /     \
 *          x ->  B       B  <- y
 *                          \
 *                           R
 */
func (t *Tree) deleteFix(x *Node) {
	var y *Node

	for x != t.root && x.color == BLACK {
		if x == x.p.left {
			y = x.p.right
			if y.color == RED {
				y.color = BLACK
				x.p.color = RED
				t.leftRotate(x.p)
				y = x.p.right
			}
			if y.right.color == BLACK && y.left.color == BLACK {
				y.color = RED
				x = x.p
			} else {
				if y.right.color == BLACK {
					y.left.color = BLACK
					y.color = RED
					t.rightRotate(y)
					y = x.p.right
				}
				y.color = x.p.color
				x.p.color = BLACK
				y.right.color = BLACK
				t.leftRotate(x.p)
				x = t.root
			}
		} else {
			y = x.p.left
			if y.color == RED {
				y.color = BLACK
				x.p.color = RED
				t.rightRotate(x.p)
				y = x.p.left
			}
			if y.left.color == BLACK && y.right.color == BLACK {
				y.color = RED
				x = x.p
			} else {
				if y.left.color == BLACK {
					y.right.color = BLACK
					y.color = RED
					t.leftRotate(y)
					y = x.p.left
				}
				y.color = x.p.color
				x.p.color = BLACK
				y.left.color = BLACK
				t.rightRotate(x.p)
				x = t.root
			}
		}
	}
	x.color = BLACK
}

// transplant s to the position of t
func (t *Tree) transplant(pos, n *Node) {
	if pos.p == t.null {
		t.root = n
	} else if pos == pos.p.left {
		pos.p.left = n
	} else {
		pos.p.right = n
	}
	n.p = pos.p
}

func (t *Tree) newNode(v interface{}) *Node {
	return &Node{
		left:  t.null,
		right: t.null,
		p:     t.null,
		v:     v,
		color: RED,
	}
}

/*
 *           x
 *          / \
 *         a   y
 *            / \
 *           b   c
 * ->
 *           y
 *          / \
 *         x   c
 *        / \
 *       a   b
 */
func (t *Tree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != t.null {
		y.left.p = x
	}
	t.transplant(x, y)
	y.left = x
	x.p = y
}

/*
 *           x
 *          / \
 *         y   c
 *        / \
 *       a   b
 * ->
 *           y
 *          / \
 *         a   x
 *            / \
 *           b   c
 */
func (t *Tree) rightRotate(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != t.null {
		y.right.p = x
	}
	t.transplant(x, y)
	y.right = x
	x.p = y
}
