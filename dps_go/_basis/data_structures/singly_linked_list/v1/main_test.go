package main

import (
	"reflect"
	"testing"
)

func Test_List_int(t *testing.T) {
	ll := NewList[int]()
	{
		ll.PushBack(100)
		ll.PushBack(101)
		ll.PushBack(-100)

		ll.PushFront(1000)
		ll.PushFront(1)
		ll.PushFront(2)
	}
	want := []int{2, 1, 1000, 100, 101, -100}
	wantSize := len(want)
	got := ll.ToSlice()

	// ***

	if ll.Size() != 6 {
		t.Errorf("want: %v, got: %v", wantSize, ll.Size())
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
