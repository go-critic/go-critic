package checker_test

import (
	"sync"
)

func notInvoked() {
	/*! possible sync.OnceFunc misuse, sync.OnceFunc(foo) result is not used */
	sync.OnceFunc(foo)
}

func notOnceAtAll() {
	/*! possible sync.OnceFunc misuse, consider to assign sync.OnceFunc(foo) to a variable */
	sync.OnceFunc(foo)()
}

func foo() {}
