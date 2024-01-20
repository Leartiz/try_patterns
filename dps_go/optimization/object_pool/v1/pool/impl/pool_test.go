package impl

import (
	"dps_go/optimization/object_pool/v1/domain"
	"dps_go/optimization/object_pool/v1/pool"
	"testing"
)

// visual?
func Test_Capture(t *testing.T) {
	pool := NewPool([]pool.Object{
		domain.NewProduct("Test", 123),
		domain.NewProduct("Test1", 223),
		domain.NewProduct("Test2", 323),
	})

	pool.println()

	object1, err := pool.CaptureObject()
	if err != nil {
		t.Error()
	}
	product1 := object1.(*domain.Product)
	product1.Name = "New name"

	pool.println()

	object2, err := pool.CaptureObject()
	if err != nil {
		t.Error()
	}
	product2 := object2.(*domain.Product)
	product2.Name = "New name 2"

	pool.println()

	pool.ReleaseObject(object1)
	pool.ReleaseObject(object2)

	pool.println()
}
