package gotine

import (
	"github.com/panjf2000/ants/v2"
	"sync"
)

// TaskSubmit submit func array
func TaskSubmit(fs []func()) {
	var g sync.WaitGroup
	g.Add(len(fs))
	for _, f := range fs {
		err := ants.Submit(funcWrapper(f, &g))
		if err != nil {
			return
		}
	}
	g.Wait()
}

func funcWrapper(f func(), g *sync.WaitGroup) func() {
	return func() {
		defer g.Done()
		f()
	}
}
