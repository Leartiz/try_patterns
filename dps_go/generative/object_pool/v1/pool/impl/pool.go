package impl

import (
	"dps_go/generative/object_pool/v1/pool"
	"fmt"
)

type objectWithStatus struct {
	object   pool.Object // <--- user object!
	captured bool
}

type Pool struct {
	objects []objectWithStatus
}

func NewPool(objects []pool.Object) *Pool {
	withStatuses := make([]objectWithStatus, len(objects))
	for i := range objects {
		withStatuses[i] = objectWithStatus{
			object:   objects[i],
			captured: false,
		}
	}

	return &Pool{
		objects: withStatuses,
	}
}

// -----------------------------------------------------------------------

func (p *Pool) CaptureObject() (pool.Object, error) {
	for i := range p.objects {
		if !p.objects[i].captured {
			p.objects[i].captured = true
			return p.objects[i].object, nil
		}
	}
	return nil, ErrAllObjectsAreCaptured
}

func (p *Pool) ReleaseObject(object pool.Object) {
	for i := range p.objects {
		if p.objects[i].object.Eq(object) {
			p.objects[i].captured = false
			return
		}
	}
}

func (p *Pool) CapturedSize() int {
	count := 0
	for i := range p.objects {
		if p.objects[i].captured {
			count++
		}
	}
	return count
}

func (p *Pool) Size() int {
	return len(p.objects)
}

// -----------------------------------------------------------------------

func (p *Pool) println() {
	for i := range p.objects {
		fmt.Println(p.objects[i].object,
			p.objects[i].captured)
	}
	fmt.Println()
}
