package checker_test

func badReturn() [](func()) {
	return nil
}

func veryBadReturn() [](func([](func()))) {
	return nil
}

var _ [](func())
var _ [5](*int)
var _ [](func())

var (
	_ int
	_ [5](*int)
	_ [](func())
)

const _ (int) = 5

type _ (int)
