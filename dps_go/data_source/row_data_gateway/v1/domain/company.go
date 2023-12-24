package domain

import "fmt"

type Company interface {
	GetId() int
	GetName() string

	SetName(value string) bool

	String() string
	Copy() Company
}

type company struct {
	id   int
	name string
}

func NewCompanyWithoutChecks(
	id int, name string,
) Company {
	return &company{
		id:   id,
		name: name,
	}
}

// getters
// -----------------------------------------------------------------------

func (c *company) GetId() int {
	return c.id
}

func (c *company) GetName() string {
	return c.name
}

func (c *company) SetName(value string) bool {
	c.name = value // <--- any
	return true
}

func (c *company) String() string {
	return fmt.Sprintf("%v", *c)
}

func (c *company) Copy() Company {
	copy := *c
	return &copy
}
