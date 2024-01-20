package main

import (
	"dps_go/optimization/object_pool/v1/domain"
	"dps_go/optimization/object_pool/v1/pool"
	"dps_go/optimization/object_pool/v1/pool/impl"
	"fmt"
)

func usage_example_0(pool pool.Pool) {
	product, _ := pool.CaptureObject()
	fmt.Println("Product:", product)
	fmt.Println("Captured size:", pool.CapturedSize())

	pool.ReleaseObject(product)
	fmt.Println("Captured size:", pool.CapturedSize())
}

func usage_example_1(pool pool.Pool) {
	object, _ := pool.CaptureObject()
	fmt.Println("Product:", object)

	product, converted := object.(*domain.Product)
	if !converted {
		fmt.Println("Object is not product")
		return
	}
	product.Name = "New name"
	pool.ReleaseObject(product)

	object, _ = pool.CaptureObject()
	fmt.Println("Product:", object)
}

func usage_example_3(poo pool.Pool) {
	//...
}

// -----------------------------------------------------------------------

func main() {
	fmt.Println("Object pool")

	products := []pool.Object{
		domain.NewProduct("Вода", 50),
		domain.NewProduct("Кофе", 30),
		//...
	}

	var pool pool.Pool = impl.NewPool(products)
	fmt.Println("Pool size:", pool.Size())

	usage_example_0(pool)
	usage_example_1(pool)
}
