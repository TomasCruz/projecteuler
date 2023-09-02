package projecteuler

import "sync"

type ConcurrentStringSet struct {
	m    map[string]struct{}
	lock sync.RWMutex
}

func NewConcurrentStringSet() ConcurrentStringSet {
	return ConcurrentStringSet{
		m:    map[string]struct{}{},
		lock: sync.RWMutex{},
	}
}

func (css *ConcurrentStringSet) Write(s string) {
	css.lock.Lock()
	defer css.lock.Unlock()
	css.m[s] = struct{}{}
}

func (css *ConcurrentStringSet) Read() []string {
	sSlice := make([]string, 0, len(css.m))
	for s := range css.m {
		sSlice = append(sSlice, s)
	}

	return sSlice
}
