package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// key - priority, value - payload
type Item[T_Key constraints.Ordered, T_Val any] struct {
	Key   T_Key
	Value T_Val
}

type PriorityQueue[T_Key constraints.Ordered, T_Val any] struct {
	items []Item[T_Key, T_Val]
}

func (p *PriorityQueue[T_Key, T_Val]) Insert(key T_Key, value T_Val) {
	p.items = append(p.items, Item[T_Key, T_Val]{
		Key:   key,
		Value: value,
	})
	p.siftUp()
}

func (p *PriorityQueue[T_Key, T_Val]) siftUp() {
	i := len(p.items) - 1
	for i != 0 {
		parent := (i - 1) / 2
		if p.items[i].Key > p.items[parent].Key {
			p.items[i], p.items[parent] = p.items[parent], p.items[i]
			i = parent
			continue
		}
		break
	}
}

func (p *PriorityQueue[T_Key, T_Val]) siftDown() {
}

func (p *PriorityQueue[T_Key, T_Val]) ExtractMaximum() (T_Key, T_Val) {
	var value T_Val
	var key T_Key
	return key, value
}

func (p *PriorityQueue[T_Key, T_Val]) GetKeys() []T_Key {
	keys := make([]T_Key, 0, len(p.items))
	for _, item := range p.items {
		keys = append(keys, item.Key)
	}
	return keys
}

func (p *PriorityQueue[T_Key, T_Val]) PrintKeys() {
	for _, item := range p.items {
		fmt.Printf("%v ", item.Key)
	}
}

func main() {
	pq := PriorityQueue[int, int]{}
	pq.Insert(20, 1)
	pq.Insert(15, 2)
	pq.Insert(10, 3)
	pq.Insert(5, 4)
	pq.Insert(25, 4)
	pq.PrintKeys()
	fmt.Println()
}
