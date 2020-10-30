// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package avl implements an AVL tree.
package avl

// Item represents a value in the tree.
type Item interface {
	// Less compares whether the current item is less than the given Item.
	Less(than Item) bool
}

// Int implements the Item interface.
type Int int

// Less implements the Item Less method.
func (a Int) Less(than Item) bool {
	b, _ := than.(Int)
	return a < b
}

// String implements the Item interface.
type String string

// Less implements the Item Less method.
func (a String) Less(than Item) bool {
	b, _ := than.(String)
	return a < b
}

// New returns a new AVL tree.
func New() *Tree {
	return &Tree{}
}

// Tree represents an AVL tree.
type Tree struct {
	length int
	root   *Node
}

// Length returns the number of items currently in the AVL tree.
func (t *Tree) Length() int {
	return t.length
}

// Root returns the root node of the AVL tree.
func (t *Tree) Root() *Node {
	return t.root
}

// Search searchs the Item of the AVL tree.
func (t *Tree) Search(item Item) Item {
	return t.root.search(item).Item()
}

// SearchNode searchs the node of the AVL tree with the item.
func (t *Tree) SearchNode(item Item) *Node {
	return t.root.search(item)
}

// Insert inserts the value v into the AVL tree.
func (t *Tree) Insert(item Item) {
	var ok bool
	t.root, ok = t.root.insert(item)
	if ok {
		t.length++
	}
}

// Delete deletes the node of the AVL tree with the value v.
func (t *Tree) Delete(item Item) {
	var ok bool
	t.root, ok = t.root.delete(item)
	if ok {
		t.length--
	}
}

// Node represents a node in the AVL tree.
type Node struct {
	height int
	left   *Node
	right  *Node
	item   Item
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

// Item returns the item of this node.
func (n *Node) Item() Item {
	if n == nil {
		return nil
	}
	return n.item
}

func (n *Node) search(item Item) *Node {
	if n != nil {
		if item.Less(n.item) {
			return n.left.search(item)
		} else if n.item.Less(item) {
			return n.right.search(item)
		} else {
			return n
		}
	}
	return nil
}

func (n *Node) insert(item Item) (root *Node, ok bool) {
	if n == nil {
		return &Node{item: item}, true
	}
	if item.Less(n.item) {
		n.left, ok = n.left.insert(item)
	} else {
		n.right, ok = n.right.insert(item)
	}
	return n.rebalance(), ok
}

func (n *Node) delete(item Item) (root *Node, ok bool) {
	if n == nil {
		return nil, false
	}
	if item.Less(n.item) {
		n.left, ok = n.left.delete(item)
		return n.rebalance(), ok
	} else if n.item.Less(item) {
		n.right, ok = n.right.delete(item)
		return n.rebalance(), ok
	} else {
		if n.left == nil && n.right == nil {
			return nil, true
		}
		if n.right == nil {
			return n.left, true
		}
		if n.left == nil {
			return n.right, true
		}
		min, right := n.right.deleteMin()
		min.right = right
		min.left = n.left
		return min.rebalance(), true
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
