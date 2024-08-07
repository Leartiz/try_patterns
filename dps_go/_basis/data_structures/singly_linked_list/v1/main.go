package main

import (
	"fmt"
)

// -----------------------------------------------------------------------

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewNode[T any](value T, next *Node[T]) *Node[T] {
	return &Node[T]{
		value: value,
		next:  next,
	}
}

// -----------------------------------------------------------------------

type List[T any] struct {
	beg, end *Node[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{
		beg: nil,
		end: nil,
	}
}

// -----------------------------------------------------------------------

func (ll *List[T]) PushBack(value T) {
	node := NewNode(value, nil)
	if ll.beg == nil {
		ll.beg = node
		ll.end = ll.beg
	} else {
		ll.end.next = node
		ll.end = ll.end.next
	}
}

func (ll *List[T]) PopBack() {
	if ll.beg == nil {
		return
	}

	if ll.beg == ll.end {
		ll.beg = nil
		ll.end = nil
		return
	}

	step := ll.beg
	for step.next != ll.end {
		step = step.next
	}
	ll.end = step
	step.next = nil
}

// -----------------------------------------------------------------------

func (ll *List[T]) PushFront(value T) {
	node := NewNode(value, nil)
	if ll.beg == nil {
		ll.beg = node
		ll.end = ll.beg
	} else {
		node.next = ll.beg
		ll.beg = node
	}
}

func (ll *List[T]) PopFront() {
	if ll.beg == nil {
		return
	}
	if ll.beg == ll.end {
		ll.beg = nil
		ll.end = nil
		return
	}

	_ = ll.beg // remove!

	ll.beg = ll.beg.next
	if ll.beg == nil {
		ll.end = nil
	}
}

// -----------------------------------------------------------------------

func (ll *List[T]) Println() {
	if ll.beg == nil {
		fmt.Println("<empty>")
		return
	}

	step := ll.beg
	for step != nil {
		fmt.Printf("%v ", step.value)
		step = step.next
	}
	fmt.Println()
}

func (ll *List[T]) Size() int {
	sz := 0
	step := ll.beg
	for step != nil {
		sz++
		step = step.next
	}
	return sz
}

func (ll *List[T]) ToSlice() []T {
	values := []T{}
	step := ll.beg
	for step != nil {
		values = append(values, step.value)
		step = step.next
	}
	return values
}

// -----------------------------------------------------------------------

func main() {

	/* visual tests */

	fmt.Println("*** int ***")
	{
		l := List[int]{}
		l.PushBack(100)
		l.PushBack(101)
		l.PushBack(102)
		l.Println()
		fmt.Printf("size: %v\n", l.Size())
		fmt.Printf("values: %v\n", l.ToSlice())

		// ***

		l.PopBack()
		l.Println()
		fmt.Printf("size: %v\n", l.Size())
		fmt.Printf("values: %v\n", l.ToSlice())

		// ***

		l.PopFront()
		l.PopFront()
		l.PopFront()
		l.PopFront()
		l.PopFront()
		l.Println()
		fmt.Printf("size: %v\n", l.Size())
		fmt.Printf("values: %v\n", l.ToSlice())

		// ***

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3)
		l.PushBack(4)
		l.PushBack(5)
		l.PushBack(6)
		l.Println()
		fmt.Printf("size: %v\n", l.Size())
		fmt.Printf("values: %v\n", l.ToSlice())
	}
	fmt.Println("*** string ***")
	{
		l := List[string]{}
		l.PushBack("abc")
		l.PushBack("ddd")
		l.PushBack("fff")
		l.Println()
		fmt.Printf("size: %v\n", l.Size())
		fmt.Printf("values: %v\n", l.ToSlice())

		// ***

		l.PopBack()
		l.Println()
		fmt.Printf("size: %v\n", l.Size())
		fmt.Printf("values: %v\n", l.ToSlice())
	}
}
