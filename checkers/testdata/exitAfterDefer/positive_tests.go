package checker_test

import (
	"os"
)

func simpleExitAfterDefer() {
	defer println("before return")
	/*! os.Exit will exit, and `defer println("before return")` will not run */
	os.Exit(0)
}

func conditionalExitAfterDefer(cond bool) {
	defer println("I'm deferred")
	if cond {
		/*! os.Exit will exit, and `defer println("I'm deferred")` will not run */
		os.Exit(0)
	}
}

func twoExits1(cond1, cond2 bool) {
	if cond1 {
		// This one is OK.
		// Nothing is deferred so far.
		os.Exit(0)
	}
	defer println("")
	if cond2 {
		/*! os.Exit will exit, and `defer println("")` will not run */
		os.Exit(0)
	}
}

func twoExits2() {
	// Only the first exit gives a warning.

	defer println("")
	/*! os.Exit will exit, and `defer println("")` will not run */
	os.Exit(0)
	os.Exit(0)
}

func deferLambda() {
	defer func(x int) {
		println(x)
	}(1)

	/*! os.Exit will exit, and `defer func(x int){...}(...)` will not run */
	os.Exit(0)
}
