package checker_test

type foo struct {
	k string
}

type bar struct{}

type iface interface {
	ptrRecv()
}

func (f foo) bar(i int) {}

func (f foo) bar2(i int, s string) {}

func (f *foo) ptrRecv() {}

func methodExprCalls() {
	f := foo{}
	/*! consider to change `foo.bar` to `f.bar` */
	foo.bar(f, 20)
	/*! consider to change `foo.bar2` to `f.bar2` */
	foo.bar2(f, 20, "str")

	/*! consider to change `(*foo).ptrRecv` to `f.ptrRecv` */
	(*foo).ptrRecv(&f)
	/*! consider to change `iface.ptrRecv` to `f.ptrRecv` */
	iface.ptrRecv(&f)

	var nilVar *foo
	/*! consider to change `(*foo).ptrRecv` to `nilVar.ptrRecv` */
	(*foo).ptrRecv(nilVar)
}
