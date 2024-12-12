package main

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
		prev:  nil,
		next:  nil,
	}
}

type List[T any] struct {
	beg *Node[T]
	end *Node[T]
}

// -----------------------------------------------------------------------
