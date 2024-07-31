package main

import (
	"fmt"
)

type CompareThat[T any] interface {
	Less(value T) bool
	More(value T) bool
	Equal(value T) bool
}

// -----------------------------------------------------------------------

type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type String string

// -----------------------------------------------------------------------

func (i Int8) More(value Int8) bool {
	return i > value
}

func (i Int8) Less(value Int8) bool {
	return i < value
}

func (i Int8) Equal(value Int8) bool {
	return i == value
}

// -----------------------------------------------------------------------

func (i Int16) More(value Int16) bool {
	return i > value
}

func (i Int16) Less(value Int16) bool {
	return i < value
}

func (i Int16) Equal(value Int16) bool {
	return i == value
}

// -----------------------------------------------------------------------

func (i Int32) More(value Int32) bool {
	return i > value
}

func (i Int32) Less(value Int32) bool {
	return i < value
}

func (i Int32) Equal(value Int32) bool {
	return i == value
}

// -----------------------------------------------------------------------

func (i Int64) More(value Int64) bool {
	return i > value
}

func (i Int64) Less(value Int64) bool {
	return i < value
}

func (i Int64) Equal(value Int64) bool {
	return i == value
}

// -----------------------------------------------------------------------

func (i String) More(value String) bool {
	return i > value
}

func (i String) Less(value String) bool {
	return i < value
}

func (i String) Equal(value String) bool {
	return i == value
}

// -----------------------------------------------------------------------

type Node[T CompareThat[T]] struct {
	value T
	l, r  *Node[T]
}

func NewNode[T CompareThat[T]](value T) *Node[T] {
	return &Node[T]{
		value: value,
		l:     nil,
		r:     nil,
	}
}

// -----------------------------------------------------------------------

type Tree[T CompareThat[T]] struct {
	root *Node[T]
}

func NewTree[T CompareThat[T]]() *Tree[T] {
	return &Tree[T]{
		root: nil,
	}
}

// -----------------------------------------------------------------------

// O(log(n))
func (t *Tree[T]) AddValue(value T) {
	node := NewNode(value)
	if t.root == nil {
		t.root = node
		return
	}

	t.addValue(t.root, value)
}

func (t *Tree[T]) addValue(node *Node[T], value T) {
	if node.value.More(value) {
		if node.l == nil {
			node.l = NewNode(value)
			return
		} else {
			t.addValue(node.l, value)
		}
	} else if node.value.Less(value) {
		if node.r == nil {
			node.r = NewNode(value)
			return
		} else {
			t.addValue(node.r, value)
		}
	}

	// else: ignore.
}

// -----------------------------------------------------------------------

// O(log(n))
func (t *Tree[T]) DelValue(value T, cond func(lhs, rhs T) bool) {
	if t.root == nil {
		return // err?
	}

	t.root = t.delValue(t.root, value, cond)
}

func (t *Tree[T]) delValue(node *Node[T], value T,
	cond func(lhs, rhs T) bool) *Node[T] {

	if node == nil {
		return node
	}

	if cond(node.value, value) {
		if node.l == nil && node.r == nil { // leaf!
			return nil
		} else {
			if node.l == nil {
				return node.r
			} else if node.r == nil {
				return node.l
			} else { // parent!
				minNode := t.Minimum(node.r)
				node.value = minNode.value

				node.r = t.delValue(node.r, minNode.value, cond) // ?
				return node
			}
		}

	} else {
		if node.value.More(value) {
			node.l = t.delValue(node.l, value, cond)
		} else {
			node.r = t.delValue(node.r, value, cond)
		}
	}
	return node
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

// O(log(n))
func (t *Tree[T]) FastSearch(desiredValue T) bool {
	return t.fastSearch(t.root, desiredValue)
}

func (t *Tree[T]) fastSearch(node *Node[T], desiredValue T) bool {
	if node == nil {
		return false
	}

	if node.value.More(desiredValue) {
		return t.fastSearch(node.l, desiredValue)
	} else if node.value.Less(desiredValue) {
		return t.fastSearch(node.r, desiredValue)
	} else {
		return true
	}
}

// -----------------------------------------------------------------------

// O(log(n))
func (t *Tree[T]) Minimum(x *Node[T]) *Node[T] {
	if x.l == nil {
		return x
	}
	return t.Minimum(x.l)
}

// O(log(n))
func (t *Tree[T]) Maximum(x *Node[T]) *Node[T] {
	if x.l == nil {
		return x
	}
	return t.Maximum(x.l)
}

// -----------------------------------------------------------------------

// O(n)
func (t *Tree[T]) Size() int {
	nodeCount := 0
	t.Traversal(InOrder, func(value T) {
		nodeCount++
	})
	return nodeCount
}

// O(n)
func (t *Tree[T]) ToSlice(order Order) []T {
	values := []T{}
	t.Traversal(order, func(value T) {
		values = append(values, value)
	})
	return values
}

// O(n)
func (t *Tree[T]) Println(order Order) {
	printValue := func(value T) { fmt.Printf("%v ", value) }
	t.Traversal(order, printValue)
	fmt.Println()
}

// -----------------------------------------------------------------------

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

	t.inOrderTraversal(node.l, action)
	action(node.value)
	t.inOrderTraversal(node.r, action)
}

func (t *Tree[T]) postOrderTraversal(node *Node[T], action func(value T)) {
	if node == nil {
		return
	}

	t.postOrderTraversal(node.l, action)
	t.postOrderTraversal(node.r, action)
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

	fmt.Println("*** Int32 ***")
	{
		t := NewTree[Int32]()
		t.AddValue(102)
		t.AddValue(100)
		t.AddValue(103)
		t.AddValue(101)
		t.AddValue(104)
		t.AddValue(103) // !

		t.Println(InOrder)
		t.Println(PreOrder)
		t.Println(PostOrder)
		fmt.Println(t.ToSlice(InOrder))
		fmt.Println(t.ToSlice(PreOrder))
		fmt.Println(t.ToSlice(PostOrder))

		// ***

		t.DelValue(102, func(lhs, rhs Int32) bool {
			return lhs == rhs
		})

		// ***

		t.Println(InOrder)
		t.Println(PreOrder)
		t.Println(PostOrder)
		fmt.Println(t.ToSlice(InOrder))
		fmt.Println(t.ToSlice(PreOrder))
		fmt.Println(t.ToSlice(PostOrder))

		// ***

		has := t.Search(100, func(lhs, rhs Int32) bool {
			return lhs.Equal(rhs)
		})
		fmt.Printf("t has 100: %v\n", has)

		has = t.Search(102, func(lhs, rhs Int32) bool {
			return lhs.Equal(rhs)
		})
		fmt.Printf("t has 102: %v\n", has)

		has = t.Search(109, func(lhs, rhs Int32) bool {
			return lhs == rhs
		})
		fmt.Printf("t has 109: %v\n", has)

		// ***

		has = t.FastSearch(100) // +
		fmt.Printf("t has 100: %v\n", has)

		has = t.FastSearch(102)
		fmt.Printf("t has 102: %v\n", has)

		has = t.FastSearch(109)
		fmt.Printf("t has 109: %v\n", has)
	}
}
