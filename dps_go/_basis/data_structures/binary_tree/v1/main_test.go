package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_Tree_v1(t *testing.T) {
	tree := NewTree[Int32]()
	tree.AddValue(50)
	tree.AddValue(100)
	tree.AddValue(25)
	tree.AddValue(125)
	tree.AddValue(112)
	tree.AddValue(12)
	tree.AddValue(14)
	tree.AddValue(28)
	tree.AddValue(99)

	// ***

	{
		want := []Int32{12, 14, 25, 28, 50, 99, 100, 112, 125}
		got := tree.ToSlice(InOrder)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}

		wantSize := 9
		gotSize := wantSize
		if gotSize != wantSize {
			t.Errorf("want: %v, got: %v", wantSize, gotSize)
		}
	}

	// ***

	{
		tree.DelValue(25, func(lhs, rhs Int32) bool { return lhs == rhs })
		tree.DelValue(14, func(lhs, rhs Int32) bool { return lhs == rhs })

		want := []Int32{12, 28, 50, 99, 100, 112, 125}
		got := tree.ToSlice(InOrder)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}

		wantSize := 7
		gotSize := wantSize
		if gotSize != wantSize {
			t.Errorf("want: %v, got: %v", wantSize, gotSize)
		}
	}

	// ***

	{
		tree.DelValue(50, func(lhs, rhs Int32) bool { return lhs == rhs })
		tree.DelValue(99, func(lhs, rhs Int32) bool { return lhs == rhs })

		want := []Int32{12, 28, 100, 112, 125}
		got := tree.ToSlice(InOrder)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}

		wantSize := 5
		gotSize := wantSize
		if gotSize != wantSize {
			t.Errorf("want: %v, got: %v", wantSize, gotSize)
		}
	}
}

// -----------------------------------------------------------------------

func genTreeAndDesiredValue() (*Tree[Int32], Int32) {
	tree := NewTree[Int32]()

	maxNodeCount := 1000
	var maxNodeValue int32 = 500

	var lastValue Int32 = 0
	for i := 0; i < maxNodeCount; i++ {
		lastValue = Int32(rand.Int31n(maxNodeValue))
		tree.AddValue(lastValue)
	}

	return tree, lastValue
}

func Benchmark_Tree_v1(b *testing.B) {

	tree, desiredValue := genTreeAndDesiredValue()

	for i := 0; i < b.N; i++ {
		tree.FastSearch(desiredValue)
	}
}

func Benchmark_Tree_v2(b *testing.B) {

	tree, desiredValue := genTreeAndDesiredValue()

	for i := 0; i < b.N; i++ {
		tree.Search(desiredValue, func(lhs, rhs Int32) bool {
			return lhs == rhs
		})
	}
}
