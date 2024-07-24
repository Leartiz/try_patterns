package main

import (
	"reflect"
	"testing"
)

func Test_Composite(t *testing.T) {
	{
		var leaf1 Component = NewLeaf("l")
		var leaf2 Component = NewLeaf("l")

		if !reflect.DeepEqual(leaf1, leaf2) {
			t.Error("leaf1 not eq leaf2")
			return
		}
	}

	// ***

	{
		var leaf1 Component = NewLeafWithNumber()
		var leaf2 Component = NewLeafWithNumber()

		var composite1 = NewCompositeWithNumber()
		composite1.Add(leaf1)
		composite1.Add(leaf2)

		var composite2 = NewCompositeWithNumber()
		composite2.Add(composite1)
		composite2.Add(leaf1)

		composite2.Display("") // visual!
	}
}
