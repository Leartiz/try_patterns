package row

import "dps_go/data_source/row_data_gateway/v1/domain"

type Company interface {
	domain.Company

	Insert() error
	Update() error
	Delete() error
	//...
}
