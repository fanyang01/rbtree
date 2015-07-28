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

type node struct {
	left, right, p *node
	v              interface{}
	color          bool
}

// RbTree is a red-black tree
type RbTree struct {
	size       int
	null, root *node
	compare    common.CompareFunc
}

// New creates an initialized tree
func New(f common.CompareFunc) *RbTree {
	null := new(node)
	null.p = null
	null.right = null
	null.left = null
	null.color = BLACK
	return &RbTree{
		size:    0,
		null:    null,
		root:    null,
		compare: f,
	}
}

// IsEmpty returns true if the tree is empty
func (t *RbTree) IsEmpty() bool {
	return t.size == 0
}

// Len returns size of t
func (t *RbTree) Len() int {
	return t.size
}

// Clean resets a tree structure to it's initial state
func (t *RbTree) Clean() *RbTree {
	t.size = 0
	t.root = t.null
	return t
}

// Has tests if v is already in t
func (t *RbTree) Has(v interface{}) bool {
	return t.search(t.root, v) != nil
}

// Search tries to find the node containing v.
// On success, it returns payload of the node and true,
// otherwise, false will be returned to indicate the node is not found.
func (t *RbTree) Search(v interface{}) (interface{}, bool) {
	x := t.search(t.root, v)
	if x == nil {
		return nil, false
	}
	return x.v, true
}

func (t *RbTree) search(r *node, v interface{}) *node {
	x := r
	for x != t.null {
		var cmp int
		if cmp = t.compare(v, x.v); cmp == 0 {
			return x
		}
		if cmp < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	return nil
}

// Insert inserts v into correct place.
// It will returns false when v is already in this tree.
func (t *RbTree) Insert(v interface{}) bool {
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
			return false
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
	return true
}

// Replace replaces payload of a node with new v.
// It returns false when the node can't be found.
func (t *RbTree) Replace(v interface{}) (interface{}, bool) {
	var before interface{}
	var ok bool
	if before, ok = t.Delete(v); !ok {
		return nil, false
	}
	if !t.Insert(v) {
		return nil, false
	}
	return before, true
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
func (t *RbTree) insertFix(x *node) {
	var y *node

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

// Delete removes v from t. It returns false when node containing v is not found.
func (t *RbTree) Delete(v interface{}) (interface{}, bool) {
	x := t.search(t.root, v)
	if x == nil {
		return nil, false
	}

	var z *node
	y := x
	color := x.color
	if x.left == t.null {
		z = x.right
		t.transplant(x, x.right)
	} else if x.right == t.null {
		z = x.left
		t.transplant(x, x.left)
	} else {
		// y is the maxium node on x's right subtree
		// it will replace x
		y = func(n *node) *node {
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
	return x.v, true
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
func (t *RbTree) deleteFix(x *node) {
	var y *node

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
func (t *RbTree) transplant(pos, n *node) {
	if pos.p == t.null {
		t.root = n
	} else if pos == pos.p.left {
		pos.p.left = n
	} else {
		pos.p.right = n
	}
	n.p = pos.p
}

func (t *RbTree) newNode(v interface{}) *node {
	return &node{
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
func (t *RbTree) leftRotate(x *node) {
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
func (t *RbTree) rightRotate(x *node) {
	y := x.left
	x.left = y.right
	if y.right != t.null {
		y.right.p = x
	}
	t.transplant(x, y)
	y.right = x
	x.p = y
}
