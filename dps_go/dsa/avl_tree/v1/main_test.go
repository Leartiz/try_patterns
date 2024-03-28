package main

import (
	"fmt"
	"testing"
)

func Test_Height(t *testing.T) {
	node := NewNode("node")
	fmt.Println(node.Height())
}
