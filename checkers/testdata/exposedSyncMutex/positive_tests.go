package checker_test

import "sync"

/*! don't embed sync.Mutex */
type App struct {
	Port int
	sync.Mutex
	Addr string
}

/*! don't embed *sync.RWMutex */
type Cacher struct {
	*sync.RWMutex
}
