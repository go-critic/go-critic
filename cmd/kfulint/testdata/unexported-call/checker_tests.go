package checker_test

type foo struct{}

func (f foo) unexported() int { return 1 }
func (f foo) Exported() int   { return 1 }

func getFoo() foo { return foo{} }

func f() {
	var fo foo
	/// getFoo().unexported() should be exported
	getFoo().unexported()
	/// fo.unexported() should be exported
	fo.unexported()
	fo.Exported()
}
