package checker_test

/*! `IN' should not be capitalized */
/*! `OUT' should not be capitalized */
func f1(IN int) (OUT int) {
	return 0
}

/*! `IN' should not be capitalized */
/*! `X' should not be capitalized */
/*! `Y' should not be capitalized */
/*! `Z' should not be capitalized */
func f2(IN, X int) (Y, Z int) {
	return 0, 0
}

type empty struct{}

/*! `IN' should not be capitalized */
/*! `OUT' should not be capitalized */
func (IN empty) method1(OUT *int) {}

/*! `PN' should not be capitalized */
func (PN empty) method2() {}

func localBody() {
	/*! `VAR1' should not be capitalized */
	/*! `VAR2' should not be capitalized */
	VAR1, VAR2 := 1, 2

	/*! `X' should not be capitalized */
	/*! `Y' should not be capitalized */
	var X, Y = VAR1, VAR2

	{
		/*! `VAR3' should not be capitalized */
		/*! `VAR4' should not be capitalized */
		VAR3, VAR4 := X, Y

		/*! `VAR5' should not be capitalized */
		VAR5, VAR3 := VAR4, VAR3
		_, _ = VAR3, VAR5

		const (
			/*! `Const1' should not be capitalized */
			Const1 = 1
			/*! `Const2' should not be capitalized */
			Const2 = 2
			/*! `Const3' should not be capitalized */
			/*! `Const4' should not be capitalized */
			/*! `Const5' should not be capitalized */
			Const3, Const4, Const5 = 3, 4, 5
		)
	}
}
