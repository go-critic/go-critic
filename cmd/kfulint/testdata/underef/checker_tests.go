package checker_tests

type t struct {
	a  int
	t2 struct {
		b int
	}
}

func (p *t) c() {
}

func sample() {
	var k *t
	///: could simplify (*k).a to k.a
	(*k).a = 5
	///: could simplify (*k).c to k.c
	(*k).c()
	///: could simplify (*k).t2 to k.t2
	(*k).t2.b = 6
}
