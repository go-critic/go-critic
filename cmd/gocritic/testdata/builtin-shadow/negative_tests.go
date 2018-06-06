package checker_test

func noWarnigs() {
	var foo struct {
		len int
		cap int
	}

	foo.len = 123
	foo.cap = 321

	foo.len, foo.cap = foo.cap, foo.len
}
