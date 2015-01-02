package bst

import "fmt"

type Tree struct {
	Key   *int
	Value interface{}
	Left  *Tree
	Right *Tree
}

func New() *Tree {
	return &Tree{}
}

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

func (t *Tree) Delete(key int) {
	if t.Key == nil {
		return
	}
	if key == *t.Key {
		// Leaf
		if t.Left.Key == nil && t.Right.Key == nil {
			t.Key = nil
		}
		// Single Child
		if t.Left.Key != nil && t.Right.Key == nil {
			t.Key = t.Left.Key
			t.Value = t.Left.Value
			t.Left = New()
		}
		if t.Left.Key == nil && t.Right.Key != nil {
			t.Key = t.Right.Key
			t.Value = t.Right.Value
			t.Right = New()
		}
		// Two Children
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

func (t *Tree) String() string {
	if t.Key != nil {
		return fmt.Sprintf("(%d %s %s)", *t.Key, t.Left, t.Right)
	} else {
		return "()"
	}
}
