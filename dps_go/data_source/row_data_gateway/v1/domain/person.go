package domain

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
