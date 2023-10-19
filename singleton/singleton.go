package singleton

import "sync"

type Singleton struct{}

var singleton *Singleton
var lock sync.Mutex

func getInstance() *Singleton {
	if singleton != nil {
		return singleton
	}

	lock.Lock()
	defer lock.Unlock()

	singleton = &Singleton{}

	return singleton
}
