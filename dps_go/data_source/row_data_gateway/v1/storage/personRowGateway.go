package storage

import (
	"dps_go/data_source/row_data_gateway/v1/domain"
	"fmt"
)

type PersonRowGateway interface {
	domain.Person

	Insert() error
	Update() error
	UpdateWithCompanyId(companyId int) error
	Delete() error
}

type personRowGateway struct {
	domain.Person
	storage Storage
}

func MakePerson(firstName, lastName string, companyId int) (PersonRowGateway, error) {
	if !domain.IsCorrectPersonName(firstName) ||
		!domain.IsCorrectPersonName(lastName) {
		return nil, fmt.Errorf("wrong names")
	}

	// ***

	storage := Instance()
	exists, err := storage.ExistsCompany(companyId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("wrong company id")
	}

	// ***

	insertedRowId, err := storage.InsertPerson(firstName, lastName, companyId)
	if err != nil {
		return nil, err
	}
	insertedPerson, err := storage.GetPersonById(insertedRowId)
	if err != nil {
		return nil, err
	}

	return &personRowGateway{
		Person:  insertedPerson,
		storage: storage,
	}, nil
}

func FindPersonById(id int) (PersonRowGateway, error) {
	storage := Instance()
	foundPerson, err := storage.GetPersonById(id)
	if err != nil {
		return nil, err
	}

	return &personRowGateway{
		Person:  foundPerson,
		storage: storage,
	}, nil
}

func FindPersonsForCompany(companyId int) ([]PersonRowGateway, error) {
	storage := Instance()
	persons, err := storage.GetPersonsByCompanyId(companyId)
	if err != nil {
		return nil, err
	}

	results := []PersonRowGateway{}
	for _, val := range persons {
		results = append(results, &personRowGateway{
			Person: val, storage: storage,
		}, nil)
	}
	return results, nil
}

// crud
// -----------------------------------------------------------------------

func (p *personRowGateway) Insert() error {
	insertedRowId, err := p.storage.InsertPerson(
		p.GetFirstName(),
		p.GetLastName(),
		p.GetCompanyId(),
	)
	if err != nil {
		return err
	}

	// ***

	p.Person = domain.NewPersonWithoutChecks(
		insertedRowId, p.GetCompanyId(),
		p.GetFirstName(),
		p.GetLastName(),
	)
	return nil
}

func (p *personRowGateway) Update() error {
	err := p.storage.UpdatePerson(
		p.GetId(), p.GetFirstName(), p.GetLastName(),
		p.GetCompanyId(),
	)
	if err != nil {
		return err
	}

	// ***

	p.Person, err = p.storage.GetPersonById(p.GetId()) // don't have to make a request!
	return err
}

func (p *personRowGateway) UpdateWithCompanyId(companyId int) error {
	exists, err := p.storage.ExistsCompany(companyId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("wrong company id")
	}

	// ***

	err = p.storage.UpdatePerson(
		p.GetId(), p.GetFirstName(), p.GetLastName(),
		companyId,
	)
	if err != nil {
		return err
	}

	// ***

	p.Person, err = p.storage.GetPersonById(p.GetId())
	return err
}

func (p *personRowGateway) Delete() error {
	return p.storage.DeletePersonById(p.GetId())
}
