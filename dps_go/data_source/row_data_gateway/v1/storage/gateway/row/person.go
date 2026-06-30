package row

import (
	"rdg_v1/domain"
)

// or active record?
type Person interface {
	domain.Person

	Insert() error
	Update() error
	UpdateWithCompanyId(companyId int) error
	Delete() error
}
