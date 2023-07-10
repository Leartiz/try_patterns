package main

import "fmt"

type I interface {
	DoConcrete()
}

// ***

type A struct {
	value string
}

func NewA(value string) *A {
	return &A{
		value: value,
	}
}

func (a *A) DoConcrete() {
	fmt.Println("--- Struct A ---")
}

// ***

type B struct {
	value float64
}

func NewB(value float64) *B {
	return &B{
		value: value,
	}
}

func (a *B) DoConcrete() {
	fmt.Println("--- Struct B ---")
}

// ***

func main() {
	fmt.Println("dps_go|behavioral|template_method|v1")

	// ***

	{
		var a I = NewA("Hi")
		a.DoConcrete()
		fmt.Println(a.(*A))
	}
	fmt.Println()
	{
		var b I = NewB(127.0)
		b.DoConcrete()
		fmt.Println(b.(*B))
	}

}
