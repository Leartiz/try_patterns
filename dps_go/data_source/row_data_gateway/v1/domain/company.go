package domain

type Company interface {
	GetId() int
	GetName() string

	SetName(value string) bool
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
	c.name = value
	return true
}
