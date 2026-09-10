package checker_test

import "runtime"

func badGOOS() {
	/*! unknown GOOS value "foobar" */
	_ = runtime.GOOS != "foobar"

	/*! unknown GOOS value "foobar" */
	_ = "foobar" == runtime.GOOS
}

func badGOOSConst() {
	const myOS = "bados"
	/*! unknown GOOS value "bados" */
	_ = runtime.GOOS == myOS
}
