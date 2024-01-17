package pool

type Object interface {
	Eq(value Object) bool
}
