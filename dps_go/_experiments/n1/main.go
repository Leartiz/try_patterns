package main

import (
	"fmt"
	"reflect"
)

type I struct {
}

type A struct {
}

func main() {
	fmt.Println(reflect.TypeOf("str").Name())
	fmt.Println(reflect.TypeOf("str").String())

	fmt.Println(reflect.TypeOf(1).Name())
	fmt.Println(reflect.TypeOf(1).String())

	// ***

	fmt.Println(reflect.TypeOf(I{}).Name())
	fmt.Println(reflect.TypeOf(I{}).String())

	fmt.Println(reflect.TypeOf(A{}).Name())
	fmt.Println(reflect.TypeOf(A{}).String())
}
