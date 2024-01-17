package main

import (
	"dps_go/generative/object_pool/v1/domain"
	"dps_go/generative/object_pool/v1/pool"
	"dps_go/generative/object_pool/v1/pool/impl"
	"fmt"
)

func main() {
	fmt.Println("Object pool")

	products := []pool.Object{
		domain.NewProduct("Вода", 50),
		domain.NewProduct("Кофе", 30),
		//...
	}

	pool := impl.NewPool(products)
	fmt.Println("Pool size:", pool.Size())

	{
		product, _ := pool.CaptureObject()
		fmt.Println("Product:", product)
		fmt.Println(pool.CapturedSize())

		pool.ReleaseObject(product)
		fmt.Println(pool.CapturedSize())
	}
}
