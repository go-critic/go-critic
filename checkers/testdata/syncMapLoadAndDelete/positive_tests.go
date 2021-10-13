package checker_test

import "sync"

func sink(args ...interface{}) {}

func _(cond bool, m, m2 *sync.Map) {
	{
		/*! use m.LoadAndDelete to perform load+delete operations atomically */
		actual, ok := m.Load("key")
		if ok {
			m.Delete("key")
			sink(actual)
		}
	}
}
