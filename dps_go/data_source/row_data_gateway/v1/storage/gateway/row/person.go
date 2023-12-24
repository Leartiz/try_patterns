package row

import (
	"dps_go/data_source/row_data_gateway/v1/domain"
)

type Person interface {
	domain.Person

	Insert() error
	Update() error
	UpdateWithCompanyId(companyId int) error
	Delete() error
}
