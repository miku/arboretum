package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	key   *int64
	left  *Node
	right *Node
}

func New() *Node {
	return &Node{}
}

func (n *Node) Insert(key int64) {
	if n.key == nil {
		n.key = &key
		n.left = &Node{}
		n.right = &Node{}
	} else {
		if key < *n.key {
			n.left.Insert(key)
		} else {
			n.right.Insert(key)
		}
	}
}

func (n *Node) String() string {
	var buf bytes.Buffer
	if n.key != nil {
		buf.WriteString(fmt.Sprintf("(%d %s %s)", *n.key, n.left, n.right))
	}
	return buf.String()
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	tree := New()
	for i := 0; i < 10; i++ {
		tree.Insert(randInt(0, 100))
	}

	fmt.Println(tree)
}
