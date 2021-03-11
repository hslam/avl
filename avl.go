// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package avl implements an AVL tree.
package avl

// Item represents a value in the tree.
type Item interface {
	// Less compares whether the current item is less than the given Item.
	Less(than Item) bool
}

// Int implements the Item interface for int.
type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Item) bool {
	return a < b.(Int)
}

// String implements the Item interface for string.
type String string

// Less returns true if string(a) < string(b).
func (a String) Less(b Item) bool {
	return a < b.(String)
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

// Max returns the max node of the AVL tree.
func (t *Tree) Max() *Node {
	return t.root.Max()
}

// Min returns the min node of the AVL tree.
func (t *Tree) Min() *Node {
	return t.root.Min()
}

// Search searches the Item of the AVL tree.
func (t *Tree) Search(item Item) Item {
	return t.search(item).Item()
}

// SearchNode searches the node of the AVL tree with the item.
func (t *Tree) SearchNode(item Item) *Node {
	return t.search(item)
}

func (t *Tree) search(item Item) *Node {
	n := t.root
	for n != nil {
		if item.Less(n.item) {
			n = n.left
		} else if n.item.Less(item) {
			n = n.right
		} else {
			return n
		}
	}
	return nil
}

// Insert inserts the item into the AVL tree.
func (t *Tree) Insert(item Item) {
	var ok bool
	t.root, ok = t.root.insert(item)
	if ok {
		t.length++
	}
}

// Clear removes all items from the AVL tree.
func (t *Tree) Clear() {
	t.root = nil
	t.length = 0
}

// Delete deletes the node of the AVL tree with the item.
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
	parent *Node
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

// Parent returns the parent node.
func (n *Node) Parent() *Node {
	if n == nil {
		return nil
	}
	return n.parent
}

// Item returns the item of this node.
func (n *Node) Item() Item {
	if n == nil {
		return nil
	}
	return n.item
}

// Max returns the max node of this node's subtree.
func (n *Node) Max() *Node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		return n.right.Max()
	}
	return n
}

// Min returns the min node of this node's subtree.
func (n *Node) Min() *Node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		return n.left.Min()
	}
	return n
}

// Last returns the last node less than this node.
func (n *Node) Last() *Node {
	if n == nil {
		return nil
	}
	if n.left != nil {
		return n.left.Max()
	}
	left := n
	p := left.parent
	for p != nil && left == p.left {
		left = p
		p = left.parent
	}
	return p
}

// Next returns the next node more than this node.
func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	if n.right != nil {
		return n.right.Min()
	}
	right := n
	p := right.parent
	for p != nil && right == p.right {
		right = p
		p = right.parent
	}
	return p
}

func (n *Node) insert(item Item) (root *Node, ok bool) {
	if n == nil {
		return &Node{item: item}, true
	}
	if item.Less(n.item) {
		n.left, ok = n.left.insert(item)
		if n.left.Height() == 0 {
			n.left.parent = n
		}
	} else if n.item.Less(item) {
		n.right, ok = n.right.insert(item)
		if n.right.Height() == 0 {
			n.right.parent = n
		}
	} else {
		n.item = item
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
		p := n.parent
		if n.right == nil {
			n.left.parent = p
			if p != nil {
				if n == p.left {
					p.left = n.left
				} else {
					p.right = n.left
				}
			}
			return n.left, true
		}
		if n.left == nil {
			n.right.parent = p
			if p != nil {
				if n == p.left {
					p.left = n.right
				} else {
					p.right = n.right
				}
			}
			return n.right, true
		}
		var min *Node
		min, n.right = n.right.deleteMin()
		n.item = min.item
		return n.rebalance(), true
	}
}

func (n *Node) deleteMin() (min *Node, parent *Node) {
	if n.left != nil {
		min, n.left = n.left.deleteMin()
		return min, n.rebalance()
	}
	if n.right != nil {
		n.right.parent = n.parent
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
	if newParent.left != nil {
		newParent.left.parent = n
	}
	p := n.parent
	if p != nil {
		if n == p.left {
			p.left = newParent
		} else {
			p.right = newParent
		}
	}
	newParent.parent = p
	n.parent = newParent
	newParent.left = n
	n.updateHeight()
	newParent.updateHeight()
	return newParent
}

func (n *Node) rotateRight() *Node {
	newParent := n.left
	n.left = newParent.right
	if newParent.right != nil {
		newParent.right.parent = n
	}
	p := n.parent
	if p != nil {
		if n == p.left {
			p.left = newParent
		} else {
			p.right = newParent
		}
	}
	newParent.parent = p
	n.parent = newParent
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
