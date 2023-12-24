package storage

import (
	"dps_go/data_source/row_data_gateway/v1/domain"
)

type Type int

const (
	MEMORY Type = iota // <--- default!
	SQL_POSTGRES
	SQL_LITE
	//...
)

type PersonStorage interface {
	GetPersons() ([]domain.Person, error)
	GetPersonById(id int) (domain.Person, error)
	GetPersonsByCompanyId(companyId int) ([]domain.Person, error)
	InsertPerson(firstName, lastName string, companyId int) (int, error)
	UpdatePerson(id int, firstName, lastName string, companyId int) error
	DeletePersonById(id int) error
	ExistsPerson(id int) (bool, error)
}

type CompanyStorage interface {
	GetCompanies() ([]domain.Company, error)
	InsertCompany(name string) (int, error)
	ExistsCompany(id int) (bool, error)
}

type Storage interface {
	PersonStorage
	CompanyStorage
	//...
}
