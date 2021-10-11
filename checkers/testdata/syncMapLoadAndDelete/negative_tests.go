package checker_test

import "sync"

func _(cond bool, m, m2 *sync.Map) {
	{
		// Condition mismatched.
		v, ok := m.Load("key")
		if ok && cond {
			m.Delete("key")
			sink(v)
		}
	}

	{
		// Maps mismatched.
		actual, ok := m.Load("key")
		if ok {
			m2.Delete("key")
			sink(actual)
		}
	}

	{
		// Keys mismatched.
		actual, ok := m.Load("key")
		if ok {
			m.Delete("key2")
			sink(actual)
		}
	}

	{
		// Return values are ignored.
		m.Load("key")
		if cond {
			m.Delete("key2")
		}
	}

	{
		v, deleted := m.LoadAndDelete("key")
		if deleted {
			sink(v)
		}
	}
}
