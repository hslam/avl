// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package avl implements an AVL tree.
package avl

import (
	"sync"
)

// CompareFunc is the compare function returns an integer comparing two values.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
type CompareFunc func(a, b interface{}) int

// New returns a new AVL tree with the CompareFunc compare.
func New(compare CompareFunc) *Tree {
	return &Tree{Compare: compare}
}

// Tree represents an AVL tree.
type Tree struct {
	mu      sync.Mutex
	root    *Node
	Compare CompareFunc
}

// Search searchs the node of the AVL tree with the value v.
func (t *Tree) Search(v interface{}) *Node {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.root.search(t.Compare, v)
}

// Insert inserts the value v into the AVL tree.
func (t *Tree) Insert(v interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root = t.root.insert(t.Compare, v)
}

// Delete deletes the node of the AVL tree with the value v.
func (t *Tree) Delete(v interface{}) {
	defer t.mu.Unlock()
	t.mu.Lock()
	t.root = t.root.delete(t.Compare, v)
}

// Root returns the root node of the AVL tree.
func (t *Tree) Root() *Node {
	defer t.mu.Unlock()
	t.mu.Lock()
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

func (n *Node) search(compare CompareFunc, v interface{}) *Node {
	if n != nil {
		switch compare(v, n.Value) {
		case -1:
			return n.left.search(compare, v)
		case 0:
			return n
		case 1:
			return n.right.search(compare, v)
		}
	}
	return nil
}

func (n *Node) insert(compare CompareFunc, v interface{}) *Node {
	if n == nil {
		return &Node{Value: v}
	}
	switch compare(v, n.Value) {
	case -1:
		n.left = n.left.insert(compare, v)
	case 0, 1:
		n.right = n.right.insert(compare, v)
	}
	return n.rebalance()
}

func (n *Node) delete(compare CompareFunc, v interface{}) *Node {
	if n == nil {
		return nil
	}
	switch compare(v, n.Value) {
	case -1:
		n.left = n.left.delete(compare, v)
		return n.rebalance()
	case 1:
		n.right = n.right.delete(compare, v)
		return n.rebalance()
	default:
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
