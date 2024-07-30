package main

type Node[T any] struct {
	value      T
	prev, next *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
		prev:  nil,
		next:  nil,
	}
}

type List[T any] struct {
	beg, end *Node[T]
}

// -----------------------------------------------------------------------
