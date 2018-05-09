package checker_test

type paramNames struct{}

func paramNamesLoud0(IN int) (OUT int) { return 0 }

func (IN paramNames) Loud1(OUT *int) {}

func paramNamesCapitalized0(X, Y int) {}

func paramNamesCapitalized1(X int) (Y, Z int) { return 0, 0 }

func (PN paramNames) Capitalized2() {}

func duplicate(a int, b int, c int) {}
