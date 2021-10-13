package checker_test

import "sync"

var cache struct {
	sync.RWMutex
	data map[string]interface{}
}
