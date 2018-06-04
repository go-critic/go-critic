package checker_test

type foo struct{}

func f() {
	var fo foo
	/// getFoo().unexported() should be exported
	getFoo().unexported()
	/// fo.unexported() should be exported
	fo.unexported()
	fo.Exported()
}
