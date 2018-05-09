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
	(*k).a = 5
	(*k).c()
	(*k).t2.b = 6
}
