package rbtree

// First returns the leftmost node in t, which is the first in-order node.
// If t is empty, it will return nil.
func (t *Tree) First() *Node {
	if t.root == nil {
		return nil
	}
	n := t.root
	for n.left != nil {
		n = n.left
	}
	return n
}

// Last returns the rightmost node in t, which is the last in-order node.
// If t is empty, it will return nil.
func (t *Tree) Last() *Node {
	if t.root == nil {
		return nil
	}
	n := t.root
	for n.right != nil {
		n = n.right
	}
	return n
}

// Next looks up the successor of n. If n is the last node, it returns nil.
func (t *Tree) Next(n *Node) *Node {
	// right subtree is not empty
	if n.right != nil {
		x := n.right
		for x.left != nil {
			x = x.left
		}
		return x
	}
	// Right subtree is empty, backward to first non-right edge
	x := n
	for x.p != nil && x.p.right == x {
		x = x.p
	}
	if x.p == nil {
		return nil
	}
	return x.p
}

// Prev looks up the presuccessor of n. If n is the first node, it returns nil.
func (t *Tree) Prev(n *Node) *Node {
	// Left subtree is not empty
	if n.left != nil {
		x := n.left
		for x.right != nil {
			x = x.right
		}
		return x
	}
	// Left subtree is empty, backward to first non-left edge
	x := n
	for x.p != nil && x.p.left == x {
		x = x.p
	}
	if x.p == nil {
		return nil
	}
	return x.p
}

// PostorderFirst looks up the first post-order node in t.
func (t *Tree) PostorderFirst() *Node {
	if t.root == nil {
		return nil
	}
	return t.PostorderFirstNode(t.root)
}

// PostorderNext looks up the post-order successor of n.
func (t *Tree) PostorderNext(n *Node) *Node {
	if n.p != nil && n == n.p.left && n.p.right != nil {
		x := n.p.right
		for x.left != nil {
			x = x.left
		}
		return x
	}
	if n.p == nil {
		return nil
	}
	return n.p
}

// PostorderFirstNode looks up the first post-order node in subtree whose root is x.
func (t *Tree) PostorderFirstNode(x *Node) *Node {
	for {
		if x.left != nil {
			x = x.left
		} else if x.right != nil {
			x = x.right
		} else {
			return x
		}
	}
}

// PreorderFirst returns the first pre-order node of t, which obviously is the root of t.
func (t *Tree) PreorderFirst() *Node { return t.root }

// PreorderNext returns the pre-order successor of x.
func (t *Tree) PreorderNext(x *Node) *Node {
	if x.left != nil {
		return x.left
	} else if x.right != nil {
		return x.right
	}
	for x.p != nil {
		if x == x.p.left && x.p.right != nil {
			return x.p.right
		}
		x = x.p
	}
	return nil
}

// PreorderLastNode looks up the last pre-order node in subtree whose root is x.
func (t *Tree) PreorderLastNode(x *Node) *Node {
	for {
		if x.right != nil {
			x = x.right
		} else if x.left != nil {
			x = x.left
		} else {
			return x
		}
	}
}
