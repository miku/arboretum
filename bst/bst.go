package bst

import "fmt"

type Node struct {
	key   *int
	left  *Node
	right *Node
}

func New() *Node {
	return &Node{}
}

func (n *Node) Insert(key int) {
	if n.key == nil {
		n.key = &key
		n.left = New()
		n.right = New()
	} else {
		if key == *n.key {
			return
		}
		if key < *n.key {
			n.left.Insert(key)
		} else {
			n.right.Insert(key)
		}
	}
}

func (n *Node) String() string {
	if n.key != nil {
		return fmt.Sprintf("(%d %s %s)", *n.key, n.left, n.right)
	} else {
		return "()"
	}
}
