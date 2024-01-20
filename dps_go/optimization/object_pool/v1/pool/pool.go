package pool

type Pool interface {
	CaptureObject() (Object, error)
	ReleaseObject(object Object)

	CapturedSize() int
	Size() int
}
