package domain

type Product struct {
	Name string
	Cost float64
}

func NewProduct(name string, cost float64) *Product {
	return &Product{
		Name: name,
		Cost: cost,
	}
}

func (p *Product) Eq(value *Product) bool {
	return p.Name == value.Name && p.Cost == value.Cost
}
