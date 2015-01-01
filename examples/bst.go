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
	fmt.Printf("%+v\n", tree) // (10 (5 () ()) (20 (15 () ()) (30 () ())))

	// Search
	var t *bst.Tree
	t = tree.Search(8)
	fmt.Println(t) // nil

	t = tree.Search(20)
	fmt.Println(t.Value) // twenty
}
