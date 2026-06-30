package impl

import (
	"rdg_v1/domain"
	"fmt"
)

type company struct {
	id   int
	name string
}

func NewCompanyWithoutChecks(
	id int, name string,
) domain.Company {
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

func (c *company) Copy() domain.Company {
	return &*c
}
