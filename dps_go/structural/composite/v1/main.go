package main

import (
	"fmt"
)

/*
Определяет интерфейс
для всех компонентов в древовидной структуре.
*/
type Component interface {
	Execute() int
	Name() string
	Display(prefix string)
}

// -----------------------------------------------------------------------

/*
Представляет отдельный компонент,
который не может содержать другие компоненты.
*/
type Leaf struct {
	name string
}

func NewLeaf(name string) *Leaf {
	return &Leaf{
		name: name,
	}
}

var leafNumber = 0

func NewLeafWithNumber() *Leaf {
	result := NewLeaf(fmt.Sprintf("l%v", leafNumber))
	leafNumber++
	return result
}

func (l *Leaf) Execute() int {
	return 1
}

func (l *Leaf) Name() string {
	return l.name
}

func (l *Leaf) Display(prefix string) {
	fmt.Printf("%v%v ", prefix, l.name)
}

// -----------------------------------------------------------------------

/*
Представляет компонент, который может содержать другие компоненты
и реализует механизм для их добавления и удаления.
*/
type Composite struct {
	children []Component
	name     string
}

func NewComposite(name string) *Composite {
	return &Composite{
		children: make([]Component, 0),
		name:     name,
	}
}

var compositeNumber = 0

func NewCompositeWithNumber() *Composite {
	result := NewComposite(fmt.Sprintf("c%v", compositeNumber))
	compositeNumber++
	return result
}

func (c *Composite) Execute() int {
	result := 0
	for i := range c.children {
		result += c.children[i].Execute()
	}
	return result
}

func (c *Composite) Name() string {
	return c.name
}

func (c *Composite) Display(prefix string) {
	fmt.Printf("\n%v| %v: ", prefix, c.name) // extra!
	for i := range c.children {
		switch c.children[i].(type) {
		case *Leaf:
			c.children[i].Display("")
		case *Composite:
			c.children[i].Display(prefix + " ")
		}
	}
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Remove(component Component) {
	var index = -1
	for ; index < len(c.children); index++ {
		if c.children[index] == component {
			break
		}
	}

	if index != -1 {
		c.children = append(c.children[:index],
			c.children[index+1:]...)
	}
}

// -----------------------------------------------------------------------

func main() { // <--- client!
	fmt.Println("composite.v1")

	// ***

	cs0 := NewCompositeWithNumber()
	cs0.Add(NewLeafWithNumber())
	cs0.Add(NewLeafWithNumber())
	cs0.Add(NewLeafWithNumber())

	cs1 := NewCompositeWithNumber()
	cs1.Add(NewLeafWithNumber())

	cs2 := NewCompositeWithNumber()
	cs2.Add(NewLeafWithNumber())
	cs2.Add(NewLeafWithNumber())

	fmt.Printf("cs0.count: %v\n", cs0.Execute())
	fmt.Printf("cs1.count: %v\n", cs1.Execute())
	fmt.Printf("cs2.count: %v\n", cs2.Execute())

	// ***

	cs3 := NewCompositeWithNumber()
	cs3.Add(NewLeafWithNumber())
	cs3.Add(NewLeafWithNumber())
	cs3.Add(cs0)
	cs3.Add(cs1)
	cs1.Add(cs2)

	fmt.Printf("cs.count: %v\n", cs2.Execute())

	// ***

	cs2.Display("")
}
