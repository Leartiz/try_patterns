package domain

type Company interface {
	GetId() int
	GetName() string

	SetName(value string) bool

	String() string
	Copy() Company
}
