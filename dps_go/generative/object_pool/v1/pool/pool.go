package pool

type Pool interface {
	CaptureObject() (any, error)
	ReleaseObject(object Object)

	CapturedSize() int
	Size() int
}
