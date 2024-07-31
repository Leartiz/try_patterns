package main

const MaxLevel = 5

type Node struct {
	value int
	nexts []*Node
}

func NewNode(value, level int) *Node {
	return &Node{
		value: value,
		nexts: make([]*Node, level),
	}
}

type List struct {
	level int
	head  *Node
}

func NewList() *List {
	return &List{
		level: 1,
		head:  nil,
	}
}

func (l *List) InsertValue(value int) {

}

// -----------------------------------------------------------------------

func main() {

}
