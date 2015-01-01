package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/miku/arboretum/bst"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	tree := bst.New()
	for i := 0; i < 5; i++ {
		tree.Insert(randInt(0, 100))
	}
	fmt.Printf("%+v\n", tree)
	// e.g. (75 (22 (5 () ()) ()) (78 () (85 () ())))
}
