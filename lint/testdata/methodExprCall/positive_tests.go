package checker_test

type foo struct {
	k string
}

func (f foo) bar(i int) {
	println(i)
}

func (f foo) bar2(i int, s string) {
	println(i)
	println(s)
}

func f1() {
	f := foo{}

	f.bar(10)

	/// consider to change `foo.bar` to `f.bar`
	foo.bar(f, 20)
}

func f2() {
	f := foo{}

	f.bar2(10, "str")

	/// consider to change `foo.bar2` to `f.bar2`
	foo.bar2(f, 20, "str")
}
