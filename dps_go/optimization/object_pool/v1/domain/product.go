package domain

import "dps_go/optimization/object_pool/v1/pool"

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

func (p *Product) Eq(value pool.Object) bool {
	other, converted := value.(*Product)
	if !converted {
		return false
	}

	return p.Name == other.Name &&
		p.Cost == other.Cost
}
