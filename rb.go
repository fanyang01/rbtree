package rbtree

// helper functions
func isRed(n *Node) bool   { return n != nil && n.color == RED }
func isBlack(n *Node) bool { return n == nil || n.color == BLACK }

func (t *Tree) propagate(x *Node) {
	for x != nil {
		old := x.Argument
		t.compute(x)
		if t.cmpArg(x.Argument, old) == 0 {
			break
		}
		x = x.p
	}
}

func (t *Tree) insertFix(x *Node) {
	var y *Node

	for x.p != nil && x.p.color == RED {
		if x.p == x.p.p.left {
			y = x.p.p.right
			if isRed(y) {
				/*
				 * [BLACK] RED (ANY)
				 *
				 * 1.
				 *              [g]
				 *             /   \
				 *            p     y
				 *           /
				 *       -> x
				 * --->
				 *               g  <-
				 *             /   \
				 *           [p]   [y]
				 *           /
				 *          x
				 */
				x.p.color = BLACK
				y.color = BLACK
				x.p.p.color = RED
				x = x.p.p
			} else {
				if x == x.p.right {
					/*
					 * 2.
					 *             [g]
					 *            /   \
					 *           p    [y]
					 *            \
					 *          -> x
					 * --->
					 *             [g]
					 *            /   \
					 *           x    [y]
					 *          /
					 *      -> p
					 */
					x = x.p
					t.leftRotate(x)
					t.compute(x)
				}
				/*
				 * 3.
				 *             [g]
				 *            /   \
				 *           p    [y]
				 *          /
				 *      -> x
				 * --->
				 *             [p]
				 *            /   \
				 *        -> x     g
				 *                  \
				 *                  [y]
				 */
				x.p.color = BLACK
				x.p.p.color = RED
				t.rightRotate(x.p.p)
				t.propagate(x.p.right)
			}
		} else {
			y = x.p.p.left
			if isRed(y) {
				x.p.color = BLACK
				y.color = BLACK
				x.p.p.color = RED
				x = x.p.p
			} else {
				if x == x.p.left {
					x = x.p
					t.rightRotate(x)
					t.compute(x)
				}
				x.p.color = BLACK
				x.p.p.color = RED
				t.leftRotate(x.p.p)
				t.propagate(x.p.left)
			}
		}
	}
	t.root.color = BLACK
}

// x can be nil, but it should be treated as a leaf.
func (t *Tree) deleteFix(p, x *Node) {
	var y *Node

	for x != t.root && isBlack(x) {
		if x == p.left {
			y = p.right
			if isRed(y) {
				/*
				 * 1.
				 *               [p]
				 *             /     \
				 *           [x]      y
				 *                   / \
				 *                 [a] [b]
				 * --->
				 *               [y]
				 *             /     \
				 *            p      [b]
				 *           / \
				 *         [x] [a] <- y
				 */
				y.color = BLACK
				p.color = RED
				t.leftRotate(p)
				t.compute(p)
				t.propagate(y)
				y = p.right
			}
			if isBlack(y.right) && isBlack(y.left) {
				/*
				 * 2.
				 *               (p)
				 *             /     \
				 *           [x]     [y]
				 *                   / \
				 *                 [a] [b]
				 * --->
				 *               (p)  <- x
				 *             /     \
				 *           [x]      y
				 *                   / \
				 *                 [a] [b]
				 */
				y.color = RED
				x, p = p, p.p
				// Don't worry :), if p is red, loop ends and it's set to black.
			} else {
				if isBlack(y.right) {
					/*
					 * 3.
					 *               (p)
					 *             /     \
					 *           [x]     [y]
					 *                   / \
					 *                  a  [b]
					 *                 /
					 *               [c]
					 * --->
					 *               (p)
					 *             /     \
					 *           [x]     [a] <- y
					 *                   / \
					 *                 [c]  y
					 *                       \
					 *                       [b]
					 */
					y.left.color = BLACK
					y.color = RED
					t.rightRotate(y)
					t.compute(y)
					y = p.right
				}
				/*
				 * 4.
				 *               (p)
				 *             /     \
				 *           [x]     [y]
				 *                   / \
				 *                 (a)  b
				 * --->
				 *               (y)
				 *             /     \
				 *           [p]     [b]
				 *           / \
				 *         [x] (a)
				 */
				y.color = p.color
				p.color = BLACK
				y.right.color = BLACK
				t.leftRotate(p)
				t.compute(p)
				t.propagate(y)
				x, p = t.root, nil
			}
		} else {
			y = p.left
			if isRed(y) {
				y.color = BLACK
				p.color = RED
				t.rightRotate(p)
				t.compute(p)
				t.propagate(y)
				y = p.left
			}
			if isBlack(y.left) && isBlack(y.right) {
				y.color = RED
				x, p = p, p.p
			} else {
				if isBlack(y.left) {
					y.right.color = BLACK
					y.color = RED
					t.leftRotate(y)
					t.compute(y)
					y = p.left
				}
				y.color = p.color
				p.color = BLACK
				y.left.color = BLACK
				t.rightRotate(p)
				t.compute(p)
				t.propagate(y)
				x, p = t.root, nil
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

// transplant s to the position of t
func (t *Tree) transplant(pos, n *Node) {
	if pos.p == nil {
		t.root = n
	} else if pos == pos.p.left {
		pos.p.left = n
	} else {
		pos.p.right = n
	}
	if n != nil {
		n.p = pos.p
	}
}

func (t *Tree) newNode(v interface{}) *Node {
	return &Node{
		left:  nil,
		right: nil,
		p:     nil,
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
	if y.left != nil {
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
	if y.right != nil {
		y.right.p = x
	}
	t.transplant(x, y)
	y.right = x
	x.p = y
}
