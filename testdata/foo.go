package foo

type paramNames struct{}

func paramNamesLoud0(IN int) (OUT int) { return 0 }

func (IN paramNames) Loud1(OUT *int) {}

func paramNamesCapitalized0(X, Y int) {}

func paramNamesCapitalized1(X int) (Y, Z int) { return 0, 0 }

func (PN paramNames) Capitalized2() {}

func badReturn() [](func()) {
	return nil
}

func veryBadReturn() [](func([](func()))) {
	return nil
}
