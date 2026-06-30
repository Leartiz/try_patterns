package main

import (
	"slices"
	"testing"
)

func Test_PQ_PrintKeys(t *testing.T) {
	pq := PriorityQueue[int, int]{}
	pq.Insert(1, 4)
	pq.Insert(2, 5)
	pq.Insert(3, 6)
	pq.PrintKeys()
}

func Test_PQ_siftUp(t *testing.T) {
	pq := PriorityQueue[int, int]{}
	pq.Insert(20, 1)
	pq.Insert(15, 2)
	pq.Insert(10, 3)
	pq.Insert(5, 4)
	pq.Insert(25, 4)

	got := pq.GetKeys()
	want := []int{25, 20, 10, 5, 15}
	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
