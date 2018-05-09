package checker_test

func badReturn() [](func()) {
	return nil
}

func veryBadReturn() [](func([](func()))) {
	return nil
}

var badVar [](func())

var badVar2 [5](*int)
