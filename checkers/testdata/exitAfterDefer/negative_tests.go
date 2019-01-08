package checker_test

import (
	"os"
)

func exitInsideLambda() {
	defer println("before return")
	// Exit inside some anonymous function is fine.
	// Though we could check that lambda is not executed
	// inside this scope, but oh well.
	_ = func() {
		os.Exit(0)
	}
}

func returnInsteadExit(cond bool) {
	defer println("I'm deferred")
	if cond {
		println("bad cond")
		return
	}
}

func exitBeforeDefer(cond1, cond2 bool) {
	if cond1 {
		// This one is OK.
		// Nothing is deferred so far.
		os.Exit(0)
	}
	defer println("")
	if cond2 {
		return
	}
}

func justDefer() {
	defer println("")
}

func noDefers() {
	println("")
}
