//go:build go1.21

package checker_test

import (
	"sync"
)

var barOnce = sync.OnceFunc(bar)

func _() {
	b := sync.OnceFunc(bar)

	// we don't care if it's actually invoked.
	_ = b
}

func bar() {}
