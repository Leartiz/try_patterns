package storage

import (
	"dps_go/data_source/row_data_gateway/v1/domain"
	"fmt"
)

type PersonRowGateway interface {
	//SetCompanyId(value int) bool
}

type personRowGateway struct {
	person  domain.Person
	storage Storage
}

func MakePerson(firstName, lastName string, companyId int) (PersonRowGateway, error) {
	if !domain.IsCorrectPersonName(firstName) || !domain.IsCorrectPersonName(lastName) {
		return &personRowGateway{}, fmt.Errorf("wrong names")
	}

	// ***

	storage := instance()
	exists, err := storage.ExistsCompany(companyId)
	if err != nil {
		return &personRowGateway{}, err
	}

	if !exists {
		return &personRowGateway{}, fmt.Errorf("wrong company id")
	}

	// ***

	insertedRowId, err := storage.InsertPerson(firstName, lastName, companyId)
	if err != nil {
		return &personRowGateway{}, err
	}
	insertedPerson, err := storage.GetPersonById(insertedRowId)
	if err != nil {
		return &personRowGateway{}, err
	}

	return &personRowGateway{
		person:  insertedPerson,
		storage: storage,
	}, nil
}

func FindPersonById(id int) (PersonRowGateway, error) {
	return personRowGateway{}, nil
}

func FindPersonsForCompany(companyId int) ([]PersonRowGateway, error) {
	return []PersonRowGateway{}, nil
}

// crud
// -----------------------------------------------------------------------

func (p *personRowGateway) Person() domain.Person {
	return p.person
}

func (p *personRowGateway) Insert() {

}

func (p *personRowGateway) Update() {

}

func (p *personRowGateway) UpdateFirstName(value string) {

}

func (p *personRowGateway) UpdateLastName(value string) {

}

func (p *personRowGateway) Delete() {

}
