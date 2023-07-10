package main

import (
	"dps_go/behavioral/template_method/v2/domain"
	"fmt"
	"reflect"
)

// ***

type MethodA struct {
}

func (m *MethodA) Invoke() error {
	fmt.Println(reflect.TypeOf(m).String(), "Invoke")
	return nil
}

type MethodB struct {
}

func (m *MethodB) Invoke() error {
	fmt.Println(reflect.TypeOf(m).String(), "Invoke")
	return nil
}

// -----------------------------------------------------------------------

// можно оставить только один такой класс,
// в него будет внедряться реализация метода
type A struct {
	domain.Base
}

func NewA(data string, method domain.IMethod) *A {
	return &A{
		domain.Base{
			Data:   data,
			Method: method,
		},
	}
}

// ***

type B struct {
	domain.Base
}

func NewB(data string, method domain.IMethod) *B {
	return &B{
		domain.Base{
			Data:   data,
			Method: method,
		},
	}
}

// -----------------------------------------------------------------------

func main() {
	fmt.Println("dps_go|behavioral|template_method|v2")

	// ***

	var a *A = NewA("Tmp", &MethodA{})
	a.Calc()

	fmt.Println()

	var b *B = NewB("Tmp", &MethodB{})
	b.Calc()

	// ***

	fmt.Println()

	var logics = make([]domain.IBaseLogic, 0)
	logics = append(logics, a)
	logics = append(logics, b)

	for i := range logics {
		if err := logics[i].Calc(); err != nil {
			fmt.Println("ILogic, Calc failed")
		} else {
			fmt.Println("ILogic, Calc succeed")
		}
	}
}
