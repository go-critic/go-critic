package checker_test

type foo struct {
	k string
}

func (f foo) bar(i int) {}

func (f foo) bar2(i int, s string) {}

func methodExprCalls() {
	f := foo{}
	/// consider to change `foo.bar` to `f.bar`
	foo.bar(f, 20)
	/// consider to change `foo.bar2` to `f.bar2`
	foo.bar2(f, 20, "str")
}
