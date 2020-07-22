package checker_test

func foo_1() {
	foo_2()
	return
}

func foo_2() int {
	func() {}()
	return 0
}

func foo_3() {
	func() {}()
}

func foo_4() {
	func() {
		foo_1()
		return
	}()
}

func foo_5() {
	defer func() {
		defer foo_1()
		foo_1()
		return
	}()
	foo1()
	return
}

func foo_6() int {
	defer func() {}()
	return foo_2()
}

func foo_7() int {
	if true {
		defer func() {}()
	}
	return 0
}

func foo_8() int {
	if true {
		foo_1()
		foo_2()
		foo_3()

		defer func() {}()
	}
	return 0
}

type sharedData struct {
	value int
}

func (*sharedData) Lock()   {}
func (*sharedData) Unlock() {}

func issue941(d *sharedData) int {
	d.Lock()
	defer d.Unlock()
	return d.value
}
