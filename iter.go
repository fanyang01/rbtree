package rbtree

// First returns the leftmost node in t, which is the first in-order node.
// If t is empty, it will return nil.
func (t *Tree) First() *Node {
	if t.root == t.null {
		return nil
	}
	n := t.root
	for n.left != t.null {
		n = n.left
	}
	return n
}

// Last returns the rightmost node in t, which is the last in-order node.
// If t is empty, it will return nil.
func (t *Tree) Last() *Node {
	if t.root == t.null {
		return nil
	}
	n := t.root
	for n.right != t.null {
		n = n.right
	}
	return n
}

// Next looks up the successor of n. If n is the last node, it returns nil.
func (t *Tree) Next(n *Node) *Node {
	// right subtree is not empty
	if n.right != t.null {
		x := n.right
		for x.left != t.null {
			x = x.left
		}
		return x
	}
	// Right subtree is empty, backward to first non-right edge
	x := n
	for x.p != t.null && x.p.right == x {
		x = x.p
	}
	if x.p == t.null {
		return nil
	}
	return x.p
}

// Prev looks up the presuccessor of n. If n is the first node, it returns nil.
func (t *Tree) Prev(n *Node) *Node {
	// Left subtree is not empty
	if n.left != t.null {
		x := n.left
		for x.right != t.null {
			x = x.right
		}
		return x
	}
	// Left subtree is empty, backward to first non-left edge
	x := n
	for x.p != t.null && x.p.left == x {
		x = x.p
	}
	if x.p == t.null {
		return nil
	}
	return x.p
}

func (t *Tree) PostorderFirst() *Node {
	if t.root == t.null {
		return nil
	}
	return t.PostorderFirstChild(t.root)
}

func (t *Tree) PostorderNext(n *Node) *Node {
	if n.p != t.null && n == n.p.left && n.p.right != t.null {
		x := n.p.right
		for x.left != t.null {
			x = x.left
		}
		return x
	}
	if n.p == t.null {
		return nil
	}
	return n.p
}

func (t *Tree) PostorderFirstChild(x *Node) *Node {
	for {
		if x.left != t.null {
			x = x.left
		} else if x.right != t.null {
			x = x.right
		} else {
			return x
		}
	}
}
