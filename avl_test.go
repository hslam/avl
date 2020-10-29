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
	tree := New()
	if r {
		for i := n - 1; i >= 0; i-- {
			tree.Insert(Int(i))
			traverse(tree.Root(), t)
		}
	} else {
		for i := 0; i < n; i++ {
			tree.Insert(Int(i))
			traverse(tree.Root(), t)
		}
	}
	testSearch(tree, j, t)
	tree.Delete(Int(j))
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
	if node := tree.Search(Int(j)); node == nil {
		t.Error("")
	} else if int(node.Item().(Int)) != j {
		t.Error("")
	}
}

func testNilNode(tree *Tree, j int, t *testing.T) {
	if node := tree.Search(Int(j)); node != nil {
		t.Error("")
	}
}

func TestRL(t *testing.T) {
	tree := New()
	tree.Insert(Int(1))
	traverse(tree.Root(), t)
	tree.Insert(Int(3))
	traverse(tree.Root(), t)
	tree.Insert(Int(2))
	traverse(tree.Root(), t)
	testSearch(tree, 1, t)
	testSearch(tree, 2, t)
	testSearch(tree, 3, t)

}

func TestLR(t *testing.T) {
	tree := New()
	tree.Insert(Int(3))
	traverse(tree.Root(), t)
	tree.Insert(Int(1))
	traverse(tree.Root(), t)
	tree.Insert(Int(2))
	traverse(tree.Root(), t)

	testSearch(tree, 1, t)
	testSearch(tree, 2, t)
	testSearch(tree, 3, t)
}

func TestDeleteRight(t *testing.T) {
	tree := New()
	for i := 0; i < 7; i++ {
		tree.Insert(Int(i))
		traverse(tree.Root(), t)
	}
	testSearch(tree, 6, t)
	testSearch(tree, 5, t)
	testSearch(tree, 4, t)

	tree.Delete(Int(6))
	traverse(tree.Root(), t)
	tree.Delete(Int(5))
	traverse(tree.Root(), t)
	tree.Delete(Int(4))
	traverse(tree.Root(), t)
	testNilNode(tree, 6, t)
	testNilNode(tree, 5, t)
	testNilNode(tree, 4, t)
}

func TestDeleteLeft(t *testing.T) {
	tree := New()
	for i := 0; i < 7; i++ {
		tree.Insert(Int(i))
		traverse(tree.Root(), t)
	}
	testSearch(tree, 0, t)
	testSearch(tree, 1, t)
	testSearch(tree, 2, t)

	tree.Delete(Int(0))
	traverse(tree.Root(), t)
	tree.Delete(Int(1))
	traverse(tree.Root(), t)
	tree.Delete(Int(2))
	traverse(tree.Root(), t)
	testNilNode(tree, 0, t)
	testNilNode(tree, 1, t)
	testNilNode(tree, 2, t)
}

func TestEmptyTree(t *testing.T) {
	tree := New()
	tree.Delete(Int(0))
	if tree.Root().Left() != nil {
		t.Error("")
	}
	if tree.Root().Right() != nil {
		t.Error("")
	}
	if tree.Root().Item() != nil {
		t.Error("")
	}
}

func BenchmarkAVL(b *testing.B) {
	tree := New()
	for i := 0; i < b.N; i++ {
		tree.Insert(Int(i))
		tree.Search(Int(i))
		tree.Delete(Int(i))
	}
}

func TestStringLess(t *testing.T) {
	a := String("a")
	b := String("b")
	if !a.Less(b) {
		t.Error("")
	}
}
