package checker_test

import "sync"

/*! don't embed sync.Mutex */
type App struct {
	Port int
	sync.Mutex
	Addr string
}

/*! don't embed *sync.Mutex */
type srv struct {
	Port int
	Addr string
	*sync.Mutex
}

/*! don't embed *sync.RWMutex */
type Cache struct {
	*sync.RWMutex
}
