package checker_test

import (
	"sync"
)

func notInvoked() {
	/*! suggestion: sync.OnceFunc(foo)() */
	sync.OnceFunc(foo)
}

func notOnceAtAll() {
	/*! possible sync.OnceFunc misuse, consider to assign sync.OnceFunc(foo) to a variable */
	sync.OnceFunc(foo)()
}

func foo() {}
