package checker_test

import "sync"

type App struct {
	Port int
	/*! don't embed sync.Mutex */
	sync.Mutex
	Addr string
}

type srv struct {
	Port int
	Addr string
	/*! don't embed *sync.Mutex */
	*sync.Mutex
}

type Cache struct {
	/*! don't embed *sync.RWMutex */
	*sync.RWMutex
}
