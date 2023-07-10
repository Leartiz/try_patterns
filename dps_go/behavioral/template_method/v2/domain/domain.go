package domain

import (
	"fmt"
	"reflect"
)

type IBaseLogic interface {
	Calc() error
}

// шаблонные метод (?)
type IMethod interface {
	Invoke() error
}

// ***

type Base struct {
	Data   string
	Method IMethod
}

// ***

func (a *Base) doBeg() {
	fmt.Println(reflect.TypeOf(a).String(), "do Beg")
}

func (a *Base) doEnd() {
	fmt.Println(reflect.TypeOf(a).String(), "do End")
}

// ***

func (a *Base) Calc() error {
	if a.Method == nil {
		return fmt.Errorf(reflect.TypeOf(a).String() + "method is nil")
	}

	// ***

	a.doBeg()

	if a.Method.Invoke() != nil {
		return fmt.Errorf(reflect.TypeOf(a).String() + "method is err")
	}

	a.doEnd()
	return nil
}
