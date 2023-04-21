package gotine

import (
	"github.com/panjf2000/ants/v2"
	"sync"
)

type Gotine struct {
	Result []DetailInfo
	g      sync.WaitGroup
	sync.RWMutex
}

type DetailInfo struct {
	Success bool
	Info    string
	Err     error
}

// TaskSubmit submit func array
func (gt *Gotine) TaskSubmit(fs []func()) {
	gt.g.Add(len(fs))
	for _, f := range fs {
		err := ants.Submit(gt.funcWrapper(f))
		if err != nil {
			return
		}
	}
	gt.g.Wait()
}

func (gt *Gotine) funcWrapper(f func()) func() {
	return func() {
		defer gt.g.Done()
		f()
	}
}
