package main

type Node struct {
	Key    string
	height int
	l      *Node
	r      *Node
}

func NewNode(key string) *Node {
	return &Node{
		Key: key,

		height: 1,
		l:      nil,
		r:      nil,
	}
}

func Height(node *Node) int {
	if node == nil {
		return 0
	}

	return node.height
}

func (n *Node) Height() int {
	return n.height
}

func (n *Node) BalanceFactor() int {
	return Height(n.r) - Height(n.l)
}
