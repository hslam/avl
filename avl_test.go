package avl

import (
	"testing"
)

func TestAVL(t *testing.T) {
	for i := 0; i < 15; i++ {
		testAVL(15, i, true, t)
		testAVL(15, i, false, t)
	}

}

func testAVL(n, j int, r bool, t *testing.T) {
	tree := New(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	if r {
		for i := n - 1; i >= 0; i-- {
			tree.Insert(i)
			traverse(tree.Root(), t)
		}
	} else {
		for i := 0; i < n; i++ {
			tree.Insert(i)
			traverse(tree.Root(), t)
		}
	}
	testSearch(tree, j, t)
	tree.Delete(j)
	traverse(tree.Root(), t)
	testNilNode(tree, j, t)
}

func traverse(node *Node, t *testing.T) {
	factor := node.balanceFactor()
	if factor > 1 || factor < -1 {
		t.Error("")
	}
	if node != nil {
		traverse(node.Left(), t)
		traverse(node.Right(), t)
	}
}

func testSearch(tree *Tree, j int, t *testing.T) {
	if node := tree.Search(j); node == nil {
		t.Error("")
	} else if node.Value.(int) != j {
		t.Error("")
	}
}

func testNilNode(tree *Tree, j int, t *testing.T) {
	if node := tree.Search(j); node != nil {
		t.Error("")
	}
}

func TestRL(t *testing.T) {
	tree := New(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	tree.Insert(1)
	traverse(tree.Root(), t)
	tree.Insert(3)
	traverse(tree.Root(), t)
	tree.Insert(2)
	traverse(tree.Root(), t)
	testSearch(tree, 1, t)
	testSearch(tree, 2, t)
	testSearch(tree, 3, t)

}

func TestLR(t *testing.T) {
	tree := New(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	tree.Insert(3)
	traverse(tree.Root(), t)
	tree.Insert(1)
	traverse(tree.Root(), t)
	tree.Insert(2)
	traverse(tree.Root(), t)

	testSearch(tree, 1, t)
	testSearch(tree, 2, t)
	testSearch(tree, 3, t)
}

func TestDeleteRight(t *testing.T) {
	tree := New(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	for i := 0; i < 7; i++ {
		tree.Insert(i)
		traverse(tree.Root(), t)
	}
	testSearch(tree, 6, t)
	testSearch(tree, 5, t)
	testSearch(tree, 4, t)

	tree.Delete(6)
	traverse(tree.Root(), t)
	tree.Delete(5)
	traverse(tree.Root(), t)
	tree.Delete(4)
	traverse(tree.Root(), t)
	testNilNode(tree, 6, t)
	testNilNode(tree, 5, t)
	testNilNode(tree, 4, t)
}

func TestDeleteLeft(t *testing.T) {
	tree := New(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	for i := 0; i < 7; i++ {
		tree.Insert(i)
		traverse(tree.Root(), t)
	}
	testSearch(tree, 0, t)
	testSearch(tree, 1, t)
	testSearch(tree, 2, t)

	tree.Delete(0)
	traverse(tree.Root(), t)
	tree.Delete(1)
	traverse(tree.Root(), t)
	tree.Delete(2)
	traverse(tree.Root(), t)
	testNilNode(tree, 0, t)
	testNilNode(tree, 1, t)
	testNilNode(tree, 2, t)
}

func TestEmptyTree(t *testing.T) {
	tree := New(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	tree.Delete(0)
	if tree.Root().Left() != nil {
		t.Error("")
	}
	if tree.Root().Right() != nil {
		t.Error("")
	}
}

func BenchmarkAVL(b *testing.B) {
	tree := New(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})
	for i := 0; i < b.N; i++ {
		tree.Insert(i)
		tree.Search(i)
		tree.Delete(i)
	}
}
