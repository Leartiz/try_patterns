package domain

import (
	"fmt"
	"unicode"
)

type Person interface {
	GetId() int
	GetFirstName() string
	GetLastName() string
	GetCompanyId() int

	SetFirstName(value string) bool
	SetLastName(value string) bool

	String() string
	Copy() Person
}

type person struct {
	id        int
	firstName string
	lastName  string
	companyId int
}

// ctors
// -----------------------------------------------------------------------

func NewPersonWithoutChecks(
	id, companyId int,
	firstName string,
	lastName string,
) Person {
	return &person{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		companyId: companyId,
	}
}

func NewEmptyPerson() Person {
	return &person{}
}

// getters
// -----------------------------------------------------------------------

func (p *person) GetId() int {
	return p.id
}

func (p *person) GetFirstName() string {
	return p.firstName
}

func (p *person) GetLastName() string {
	return p.lastName
}

func (p *person) GetCompanyId() int {
	return p.companyId
}

func (p *person) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *person) Copy() Person {
	copyPerson := *p
	return &copyPerson
}

// setters
// -----------------------------------------------------------------------

func (p *person) SetFirstName(value string) bool {
	if IsCorrectPersonName(value) {
		p.firstName = value
		return true
	}
	return false
}

func (p *person) SetLastName(value string) bool {
	if IsCorrectPersonName(value) {
		p.lastName = value
		return true
	}
	return false
}

// validators
// -----------------------------------------------------------------------

func IsCorrectPersonName(value string) bool {
	runes := []rune(value)
	if len(runes) < 3 {
		return false
	}
	if !unicode.IsUpper(runes[0]) {
		return false
	}

	runes = runes[1:]
	for _, element := range runes {
		if !unicode.IsLower(rune(element)) {
			return false
		}
	}
	return true
}
