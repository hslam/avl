// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package avl implements an AVL tree.
package avl

// LessFunc is the less function returns an boolean.
// The result will be true if a < b.
type LessFunc func(a, b interface{}) bool

// New returns a new AVL tree with the LessFunc less.
func New(less LessFunc) *Tree {
	return &Tree{Less: less}
}

// Tree represents an AVL tree.
type Tree struct {
	root *Node
	Less LessFunc
}

// Search searchs the node of the AVL tree with the value v.
func (t *Tree) Search(v interface{}) *Node {
	return t.root.search(t.Less, v)
}

// Insert inserts the value v into the AVL tree.
func (t *Tree) Insert(v interface{}) {
	t.root = t.root.insert(t.Less, v)
}

// Delete deletes the node of the AVL tree with the value v.
func (t *Tree) Delete(v interface{}) {
	t.root = t.root.delete(t.Less, v)
}

// Root returns the root node of the AVL tree.
func (t *Tree) Root() *Node {
	return t.root
}

// Node represents a node in the AVL tree.
type Node struct {
	height int
	left   *Node
	right  *Node
	Value  interface{}
}

// Height returns the height of this node's sub-tree.
func (n *Node) Height() int {
	if n == nil {
		return -1
	}
	return n.height
}

// Left returns the left child node.
func (n *Node) Left() *Node {
	if n == nil {
		return nil
	}
	return n.left
}

// Right returns the right child node.
func (n *Node) Right() *Node {
	if n == nil {
		return nil
	}
	return n.right
}

func (n *Node) search(less LessFunc, v interface{}) *Node {
	if n != nil {
		if less(v, n.Value) {
			return n.left.search(less, v)
		} else if less(n.Value, v) {
			return n.right.search(less, v)
		} else {
			return n
		}
	}
	return nil
}

func (n *Node) insert(less LessFunc, v interface{}) *Node {
	if n == nil {
		return &Node{Value: v}
	}
	if less(v, n.Value) {
		n.left = n.left.insert(less, v)
	} else {
		n.right = n.right.insert(less, v)
	}
	return n.rebalance()
}

func (n *Node) delete(less LessFunc, v interface{}) *Node {
	if n == nil {
		return nil
	}
	if less(v, n.Value) {
		n.left = n.left.delete(less, v)
		return n.rebalance()
	} else if less(n.Value, v) {
		n.right = n.right.delete(less, v)
		return n.rebalance()
	} else {
		if n.left == nil && n.right == nil {
			return nil
		}
		if n.right == nil {
			return n.left
		}
		if n.left == nil {
			return n.right
		}
		min, right := n.right.deleteMin()
		min.right = right
		min.left = n.left
		return min.rebalance()
	}

}

func (n *Node) deleteMin() (min *Node, parent *Node) {
	if n.left != nil {
		min, n.left = n.left.deleteMin()
		return min, n.rebalance()
	}
	return n, n.right
}

func (n *Node) rebalance() *Node {
	n.updateHeight()
	balanceFactor := n.balanceFactor()
	if balanceFactor > 1 {
		if n.right.balanceFactor() < 0 {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor < -1 {
		if n.left.balanceFactor() > 0 {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

func (n *Node) updateHeight() {
	n.height = max(n.left.Height(), n.right.Height()) + 1
}

func (n *Node) balanceFactor() int {
	if n == nil {
		return 0
	}
	return n.right.Height() - n.left.Height()
}

func (n *Node) rotateLeft() *Node {
	newParent := n.right
	n.right = newParent.left
	newParent.left = n
	n.updateHeight()
	newParent.updateHeight()
	return newParent
}

func (n *Node) rotateRight() *Node {
	newParent := n.left
	n.left = newParent.right
	newParent.right = n
	n.updateHeight()
	newParent.updateHeight()
	return newParent
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
