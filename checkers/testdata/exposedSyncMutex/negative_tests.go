package checker_test

import "sync"

var cache struct {
	sync.RWMutex
	data map[string]interface{}
}

// ok for unexported
type srv struct {
	Port int
	Addr string
	*sync.Mutex
}

var Cache struct {
	sync.RWMutex
	data map[string]interface{}
}

type Global struct {
	sync.Locker
}
