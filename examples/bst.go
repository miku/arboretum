package main

import (
	"fmt"

	"github.com/miku/arboretum/bst"
)

func main() {
	tree := bst.New()

	// Insertion
	tree.Insert(10, "ten")
	tree.Insert(20, "twenty")
	tree.Insert(30, "thirty")
	tree.Insert(5, "five")
	tree.Insert(15, "fiveteen")
	fmt.Println(tree) // (10 (5 () ()) (20 (15 () ()) (30 () ())))

	// Search
	var t *bst.Tree
	t = tree.Search(8)
	fmt.Println(t) // nil

	t = tree.Search(20)
	fmt.Println(t.Value) // twenty

	// Delete
	tree.Delete(15)
	fmt.Println(tree) // (10 (5 () ()) (20 () (30 () ())))

	tree.Delete(10)
	fmt.Println(tree) // (20 (5 () ()) (30 () ()))

	tree.Delete(7)
	fmt.Println(tree) // (20 (5 () ()) (30 () ()))

	tree.Delete(30)
	fmt.Println(tree) // (20 (5 () ()) ())

	tree.Delete(5)
	fmt.Println(tree) // (20 () ())

	tree.Delete(20)
	fmt.Println(tree) // ()
}
