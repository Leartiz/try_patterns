package storage

import (
	"dps_go/data_source/row_data_gateway/v1/domain"
	"fmt"
	"sync"
)

type Storage interface {
	GetPersons() ([]domain.Person, error)
	GetPersonById(id int) (domain.Person, error)
	GetPersonsByCompanyId(companyId int) ([]domain.Person, error)
	InsertPerson(firstName, lastName string, companyId int) (int, error)
	UpdatePerson(id int, firstName, lastName string, companyId int) error
	DeletePersonById(id int) error
	ExistsPerson(id int) (bool, error)

	GetCompanies() ([]domain.Company, error)
	InsertCompany(name string) (int, error)
	ExistsCompany(id int) (bool, error)
	//...
}

type storage struct {
	rwMx      sync.RWMutex
	persons   map[int]domain.Person
	companies map[int]domain.Company

	nextPersonId  int
	nextCompanyId int
}

var once sync.Once
var storageInstance Storage

func Instance() Storage {
	once.Do(func() {
		storageInstance = &storage{
			rwMx:          sync.RWMutex{},
			persons:       make(map[int]domain.Person),
			companies:     make(map[int]domain.Company),
			nextPersonId:  0,
			nextCompanyId: 0,
		}
	})
	return storageInstance
}

// here is sql?
// -----------------------------------------------------------------------

func (s *storage) GetPersons() ([]domain.Person, error) {
	s.rwMx.RLock()
	defer s.rwMx.RUnlock()

	persons := []domain.Person{}
	for _, val := range s.persons {
		persons = append(persons, val.Copy())
	}
	return persons, nil
}

func (s *storage) GetPersonById(id int) (domain.Person, error) {
	s.rwMx.RLock()
	defer s.rwMx.RUnlock()

	person, exists := s.persons[id] // person maybe nil!
	if !exists {
		return nil, fmt.Errorf("person with id %v does not found", id)
	}

	return person.Copy(), nil
}

func (s *storage) GetPersonsByCompanyId(companyId int) ([]domain.Person, error) {
	s.rwMx.RLock()
	defer s.rwMx.RUnlock()

	persons := []domain.Person{}
	for _, val := range s.persons {
		if val.GetCompanyId() == companyId {
			persons = append(persons, val.Copy())
		}
	}
	return persons, nil
}

func (s *storage) InsertPerson(firstName, lastName string, companyId int) (int, error) {
	s.rwMx.Lock()
	defer s.rwMx.Unlock()

	// *** constraints ***
	_, exists := s.companies[companyId]
	if !exists {
		return 0, fmt.Errorf("company with id %v does not exist", companyId)
	}

	currentId := s.nextPersonId
	s.nextPersonId++

	// names checked above!

	s.persons[currentId] = domain.NewPersonWithoutChecks(
		currentId, companyId, firstName, lastName)
	return currentId, nil
}

func (s *storage) UpdatePerson(id int,
	firstName, lastName string, companyId int,
) error {
	s.rwMx.Lock()
	defer s.rwMx.Unlock()

	_, exists := s.persons[id]
	if !exists {
		return fmt.Errorf("person with id %v does not found", id)
	}

	// *** constraints ***
	_, exists = s.companies[companyId]
	if !exists {
		return fmt.Errorf("company with id %v does not exist", companyId)
	}

	s.persons[id] = domain.NewPersonWithoutChecks(
		id, companyId, firstName, lastName)
	return nil
}

func (s *storage) DeletePersonById(id int) error {
	s.rwMx.Lock()
	defer s.rwMx.Unlock()

	_, exists := s.persons[id] // person maybe nil!
	if !exists {
		return fmt.Errorf("person with id %v does not found", id)
	}
	delete(s.persons, id)
	return nil
}

func (s *storage) ExistsPerson(id int) (bool, error) {
	s.rwMx.RLock()
	defer s.rwMx.RUnlock()

	_, exists := s.persons[id]
	return exists, nil
}

// -----------------------------------------------------------------------

func (s *storage) GetCompanies() ([]domain.Company, error) {
	s.rwMx.RLock()
	defer s.rwMx.RUnlock()

	companies := []domain.Company{}
	for _, val := range s.companies {
		companies = append(companies, val.Copy())
	}
	return companies, nil
}

func (s *storage) InsertCompany(name string) (int, error) {
	s.rwMx.Lock()
	defer s.rwMx.Unlock()

	currentId := s.nextCompanyId
	s.nextCompanyId++

	s.companies[currentId] = domain.NewCompanyWithoutChecks(
		currentId, name)
	return currentId, nil
}

func (s *storage) ExistsCompany(id int) (bool, error) {
	s.rwMx.RLock()
	defer s.rwMx.RUnlock()

	_, exists := s.companies[id]
	return exists, nil
}
