package row

import "rdg_v1/domain"

type Company interface {
	domain.Company

	Insert() error
	Update() error
	Delete() error
	//...
}
