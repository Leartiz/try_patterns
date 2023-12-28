package storage

type ActiveRecord interface {
	Create() error
	Update() error
	Delete() error
}
