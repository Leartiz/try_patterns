package main

import (
	"fmt"

	"main_sequence/balanced"
	"main_sequence/pain"
	"main_sequence/useless"
)

func main() {
	fmt.Println("=== zone of pain (concrete god object) ===")
	painSvc := &pain.OrderService{}
	if err := painSvc.PlaceOrder("book-42", "user@example.com"); err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println("=== zone of uselessness (abstract factory maze) ===")
	if err := useless.PlaceOrder("book-42"); err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println("=== balanced (interface + simple service) ===")
	svc := balanced.NewService(balanced.NewMemoryRepo())
	if err := svc.PlaceOrder("book-42"); err != nil {
		fmt.Println(err)
	}
}
