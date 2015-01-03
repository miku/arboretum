// bst, a simple binary search tree implementation
package bst

import "fmt"

// Tree holds a node in a binary tree. An empty node has a nil key,
// a leaf node has nil left and right trees.
type Tree struct {
	Key   *int
	Value interface{}
	Left  *Tree
	Right *Tree
}

// New returns an empty tree
func New() *Tree {
	return &Tree{}
}

// Insert inserts key and value only if key does not exist yet
func (t *Tree) Insert(key int, value interface{}) {
	if t.Key == nil {
		t.Key = &key
		t.Value = value
		t.Left = New()
		t.Right = New()
	} else {
		if key == *t.Key {
			return
		}
		if key < *t.Key {
			t.Left.Insert(key, value)
		} else {
			t.Right.Insert(key, value)
		}
	}
}

// Search returns the Tree node if the key is found, nil otherwise
func (t *Tree) Search(key int) *Tree {
	if t.Key == nil {
		return nil
	}
	if *t.Key == key {
		return t
	}
	if key < *t.Key {
		return t.Left.Search(key)
	} else {
		return t.Right.Search(key)
	}
}

func (t *Tree) leftmostLeaf() *Tree {
	for {
		if t.Left.Key == nil {
			break
		}
		return t.Left.leftmostLeaf()
	}
	return t
}

// Delete remove the node with the given key
func (t *Tree) Delete(key int) {
	if t.Key == nil {
		return
	}
	if key == *t.Key {
		// leaf
		if t.Left.Key == nil && t.Right.Key == nil {
			t.Key = nil
		}
		// single left child
		if t.Left.Key != nil && t.Right.Key == nil {
			t.Key = t.Left.Key
			t.Value = t.Left.Value
			t.Left = New()
		}
		// single right child
		if t.Left.Key == nil && t.Right.Key != nil {
			t.Key = t.Right.Key
			t.Value = t.Right.Value
			t.Right = New()
		}
		// two children
		if t.Left.Key != nil && t.Right.Key != nil {
			successor := t.Right.leftmostLeaf()
			t.Key = successor.Key
			successor.Delete(*successor.Key)
		}
	} else {
		if key < *t.Key {
			t.Left.Delete(key)
		} else {
			t.Right.Delete(key)
		}
	}
}

// String renders the tree as abbreviated s-expressions
func (t *Tree) String() string {
	if t.Key != nil {
		return fmt.Sprintf("(%d %s %s)", *t.Key, t.Left, t.Right)
	} else {
		return "()"
	}
}
