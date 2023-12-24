package impl

import (
	domain "dps_go/data_source/row_data_gateway/v1/domain"
	domainImpl "dps_go/data_source/row_data_gateway/v1/domain/impl"
	storage "dps_go/data_source/row_data_gateway/v1/storage"
	rowGateway "dps_go/data_source/row_data_gateway/v1/storage/gateway/row"
	storageImpl "dps_go/data_source/row_data_gateway/v1/storage/impl"
	"fmt"
)

type personRowGateway struct {
	domain.Person
	storageInstance storage.Storage
}

func MakePerson(storageType storage.Type,
	firstName, lastName string, companyId int,
) (rowGateway.Person, error) {

	if !domainImpl.IsCorrectPersonName(firstName) ||
		!domainImpl.IsCorrectPersonName(lastName) {
		return nil, fmt.Errorf("wrong names")
	}

	// ***

	storage := storageImpl.Instance(storageType)
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
		Person:          insertedPerson,
		storageInstance: storage,
	}, nil
}

func FindPersonById(storageType storage.Type, id int) (rowGateway.Person, error) {
	storage := storageImpl.Instance(storageType)
	foundPerson, err := storage.GetPersonById(id)
	if err != nil {
		return nil, err
	}

	return &personRowGateway{
		Person:          foundPerson,
		storageInstance: storage,
	}, nil
}

func FindPersonsForCompany(storageType storage.Type, companyId int) ([]rowGateway.Person, error) {
	storage := storageImpl.Instance(storageType)
	persons, err := storage.GetPersonsByCompanyId(companyId)
	if err != nil {
		return nil, err
	}

	results := []rowGateway.Person{}
	for _, val := range persons {
		results = append(results, &personRowGateway{
			Person: val, storageInstance: storage,
		}, nil)
	}
	return results, nil
}

// crud
// -----------------------------------------------------------------------

func (p *personRowGateway) Insert() error {
	insertedRowId, err := p.storageInstance.InsertPerson(
		p.GetFirstName(),
		p.GetLastName(),
		p.GetCompanyId(),
	)
	if err != nil {
		return err
	}

	// ***

	p.Person = domainImpl.NewPersonWithoutChecks(
		insertedRowId, p.GetCompanyId(),
		p.GetFirstName(),
		p.GetLastName(),
	)
	return nil
}

func (p *personRowGateway) Update() error {
	err := p.storageInstance.UpdatePerson(
		p.GetId(), p.GetFirstName(), p.GetLastName(),
		p.GetCompanyId(),
	)
	if err != nil {
		return err
	}

	// ***

	p.Person, err = p.storageInstance.GetPersonById(p.GetId()) // don't have to make a request!
	return err
}

func (p *personRowGateway) UpdateWithCompanyId(companyId int) error {
	exists, err := p.storageInstance.ExistsCompany(companyId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("wrong company id")
	}

	// ***

	err = p.storageInstance.UpdatePerson(
		p.GetId(), p.GetFirstName(), p.GetLastName(),
		companyId,
	)
	if err != nil {
		return err
	}

	// ***

	p.Person, err = p.storageInstance.GetPersonById(p.GetId())
	return err
}

func (p *personRowGateway) Delete() error {
	return p.storageInstance.DeletePersonById(p.GetId())
}
