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
			testTraversal(tree, t)
		}
	} else {
		for i := 0; i < n; i++ {
			tree.Insert(Int(i))
			testTraversal(tree, t)
		}
	}
	if tree.Length() != n {
		t.Error("")
	}
	tree.Delete(Int(n))
	if tree.Length() != n {
		t.Error("")
	}
	testSearch(tree, j, t)
	tree.Delete(Int(j))
	testTraversal(tree, t)
	testNilNode(tree, j, t)
	if tree.Length() != n-1 {
		t.Error("")
	}
}

func testTraversal(tree *Tree, t *testing.T) {
	traverse(tree.Root(), t)
	testIteratorAscend(tree, t)
	testIteratorDescend(tree, t)
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

func testIteratorAscend(tree *Tree, t *testing.T) {
	iter := tree.Root().Min()
	next := iter.Next()
	for iter != nil && next != nil {
		if !iter.Item().Less(next.Item()) {
			t.Error("")
		}
		iter = next
		next = iter.Next()
	}
}

func testIteratorDescend(tree *Tree, t *testing.T) {
	iter := tree.Root().Max()
	last := iter.Last()
	for iter != nil && last != nil {
		if !last.Item().Less(iter.Item()) {
			t.Error("")
		}
		iter = last
		last = iter.Last()
	}
}

func testSearch(tree *Tree, j int, t *testing.T) {
	if node := tree.SearchNode(Int(j)); node == nil {
		t.Error("")
	} else if int(node.Item().(Int)) != j {
		t.Error("")
	}
	if item := tree.Search(Int(j)); item == nil {
		t.Error("")
	} else if int(item.(Int)) != j {
		t.Error("")
	}
}

func testNilNode(tree *Tree, j int, t *testing.T) {
	if item := tree.Search(Int(j)); item != nil {
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
	if tree.Root() != nil {
		t.Error("")
	}
	if tree.Min() != nil {
		t.Error("")
	}
	if tree.Max() != nil {
		t.Error("")
	}
	if tree.Root().Left() != nil {
		t.Error("")
	}
	if tree.Root().Right() != nil {
		t.Error("")
	}
	if tree.Root().Parent() != nil {
		t.Error("")
	}
	if tree.Root().Item() != nil {
		t.Error("")
	}
	if tree.Root().Last() != nil {
		t.Error("")
	}
	if tree.Root().Next() != nil {
		t.Error("")
	}
	if tree.Length() != 0 {
		t.Error("")
	}
	tree.Insert(Int(1))
	tree.Insert(Int(2))
	tree.Insert(Int(3))
	if tree.Root().Min().Parent() != tree.Root() {
		t.Error("")
	}
	tree.Clear()
	if tree.Length() != 0 {
		t.Error("")
	}
}

func BenchmarkAVL(b *testing.B) {
	tree := New()
	for i := 0; i < b.N; i++ {
		tree.Insert(Int(i))
		//tree.Search(Int(i))
		//tree.Delete(Int(i))
	}
}

func TestStringLess(t *testing.T) {
	a := String("a")
	b := String("b")
	if !a.Less(b) {
		t.Error("")
	}
}
