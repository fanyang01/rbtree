package rbtree

// Visitor walks in a tree. After visiting a node,
// the result visitor w is used for next node.
type Visitor interface {
	// Visit is invoked for nodes encountered by walker.
	Visit(n *Node) (w Visitor)
}

// VisitFunc is invoked for nodes. If it returns false, tree traversal will stop.
type VisitFunc func(n *Node) bool

// Visit implements the Visitor interface
func (f WalkerFunc) Visit(n *Node) (w Visitor) {
	if ok := f(n); ok {
		return f
	}
	return nil
}

// Walk traverses t in in-order, which is also ascend order of values.
func (t *Tree) Walk(v Visitor) {
	for x := t.First(); x != nil; x = t.Next(x) {
		v = v.Visit(x)
		if v == nil {
			return
		}
	}
}

// WalkReverse traverses t in descend order of values.
func (t *Tree) WalkReverse(v Visitor) {
	for x := t.Last(); x != nil; x = t.Prev(x) {
		v = v.Visit(x)
		if v == nil {
			return
		}
	}
}

// WalkPostorder traverses t in post-order, which means that a node is encountered after its children.
func (t *Tree) WalkPostorder(v Visitor) {
	for x := t.PostorderFirst(); x != nil; x = t.PostorderNext(x) {
		v = v.Visit(x)
		if v == nil {
			return
		}
	}
}

// WalkSubPostorder traverses subtree rooted at x in post-order, x self is also visited.
func (t *Tree) WalkSubPostorder(v Visitor, x *Node) {
	for n := t.PostorderFirstNode(x); n != x; n = t.PostorderNext(n) {
		v = v.Visit(n)
		if v == nil {
			return
		}
	}
	v.Visit(x)
}

// WalkPreorder traverses t in pre-order, which means that a node is encountered before its children.
func (t *Tree) WalkPreorder(v Visitor) {
	for x := t.PreorderFirst(); x != nil; x = t.PreorderNext(x) {
		v = v.Visit(x)
		if v == nil {
			return
		}
	}
}

// WalkSubPreorder traverses subtree rooted at x in pre-order, x self is also visited.
func (t *Tree) WalkSubPreorder(v Visitor, x *Node) {
	for n := t.PreorderLastNode(x); x != n; x = t.PreorderNext(n) {
		v = v.Visit(n)
		if v == nil {
			return
		}
	}
	v.Visit(n)
}
