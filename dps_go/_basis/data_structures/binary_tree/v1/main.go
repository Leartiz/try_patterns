package main

import (
	"fmt"
)

type LessThat[T any] interface {
	Less(value T) bool
}

// -----------------------------------------------------------------------

type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type String string

func (i Int8) Less(value Int8) bool {
	return i < value
}

func (i Int16) Less(value Int16) bool {
	return i < value
}

func (i Int32) Less(value Int32) bool {
	return i < value
}

func (i Int64) Less(value Int64) bool {
	return i < value
}

func (i String) Less(value String) bool {
	return i < value
}

// -----------------------------------------------------------------------

type Node[T LessThat[T]] struct {
	value T
	l, r  *Node[T]
}

func NewNode[T LessThat[T]](value T) *Node[T] {
	return &Node[T]{
		value: value,
		l:     nil,
		r:     nil,
	}
}

// -----------------------------------------------------------------------

type Tree[T LessThat[T]] struct {
	root *Node[T]
}

func NewTree[T LessThat[T]]() *Tree[T] {
	return &Tree[T]{
		root: nil,
	}
}

// -----------------------------------------------------------------------

func (t *Tree[T]) AddValue(value T) {
	node := NewNode(value)
	if t.root == nil {
		t.root = node
		return
	}

	t.addValue(t.root, value)
}

func (t *Tree[T]) addValue(node *Node[T], value T) {
	if node.value.Less(value) {
		if node.l == nil {
			node.l = NewNode(value)
			return
		} else {
			t.addValue(node.l, value)
		}
	} else {
		if node.r == nil {
			node.r = NewNode(value)
			return
		} else {
			t.addValue(node.r, value)
		}
	}
}

// -----------------------------------------------------------------------

func (t *Tree[T]) DelValue(value T, cond func(lhs, rhs T) bool) {
	if t.root == nil {
		return // err?
	}
	if cond(t.root.value, value) {
		t.root = nil
		return
	}

	t.delValue(t.root, value, cond)
}

func (t *Tree[T]) delValue(prevNode *Node[T], value T, cond func(lhs, rhs T) bool) {
	left := prevNode.l
	right := prevNode.r

	// ***

	if left != nil {
		if cond(left.value, value) {
			t.delNode(prevNode, left)
		} else {
			t.delValue(left, value, cond)
		}
	}
	if right != nil {
		if cond(right.value, value) {
			t.delNode(prevNode, right)
		} else {
			t.delValue(right, value, cond)
		}
	}
}

func (t *Tree[T]) delNode(prevNode, targetNode *Node[T]) {

}

// -----------------------------------------------------------------------

type Order int

const (
	InOrder Order = iota
	PreOrder
	PostOrder
)

// O(n)
func (t *Tree[T]) Search(desiredValue T, cond func(lhs, rhs T) bool) bool {
	result := false
	searchValue := func(value T) {
		if cond(value, desiredValue) {
			result = true
		}
	}
	t.Traversal(InOrder, searchValue)
	return result
}

// -----------------------------------------------------------------------

// O(n)
func (t *Tree[T]) Println(order Order) {
	printValue := func(value T) { fmt.Printf("%v ", value) }
	t.Traversal(order, printValue)
	fmt.Println()
}

func (t *Tree[T]) preOrderTraversal(node *Node[T], action func(value T)) {
	if node == nil {
		return
	}

	action(node.value)
	t.preOrderTraversal(node.l, action)
	t.preOrderTraversal(node.r, action)
}

func (t *Tree[T]) inOrderTraversal(node *Node[T], action func(desiredValue T)) {
	if node == nil {
		return
	}

	t.preOrderTraversal(node.l, action)
	action(node.value)
	t.preOrderTraversal(node.r, action)
}

func (t *Tree[T]) postOrderTraversal(node *Node[T], action func(value T)) {
	if node == nil {
		return
	}

	t.preOrderTraversal(node.l, action)
	t.preOrderTraversal(node.r, action)
	action(node.value)
}

// O(n)
func (t *Tree[T]) Traversal(order Order, action func(value T)) {
	switch order {
	case PreOrder:
		t.preOrderTraversal(t.root, action)
	case InOrder:
		t.inOrderTraversal(t.root, action)
	case PostOrder:
		t.postOrderTraversal(t.root, action)
	}
}

// -----------------------------------------------------------------------

func main() {
	{
		var a int = 101
		fmt.Printf("a: %v\n", a)
	}
	fmt.Println("*** Int32 ***")
	{
		t := NewTree[Int32]()
		t.AddValue(100)
		t.AddValue(101)
		t.AddValue(102)
		t.AddValue(103)
		t.AddValue(104)

		// ***

		t.Println(InOrder)
		t.Println(PreOrder)
		t.Println(PostOrder)

		// ***

		has := t.Search(102, func(lhs, rhs Int32) bool {
			return lhs == rhs
		})
		fmt.Printf("t has 102: %v\n", has)

		has = t.Search(109, func(lhs, rhs Int32) bool {
			return lhs == rhs
		})
		fmt.Printf("t has 109: %v\n", has)
	}
}
