package checker_test

func subExpr() {
	o := newObj()

	_ = o.a.c.d.e[0].p1 + o.a.c1.d.e[0].p1 + o.a.c2.d.e[0].p1

	_ = o.a.c.d.e[0].p1 + o.a.c1.d.e[0].p2 + o.a.c2.d.e[0].p3

	_ = (o.a.c.x + o.a.c.x) +
		(o.a.c1.y + o.a.c1.y) +
		(o.a.c2.z + o.a.c2.z)
}

func caseExpr() {
	o := newObj()

	switch {
	case o.a.c.d.e[0].p1 == 0:
	case o.a.c1.d.e[0].p1 == 1:
	case o.a.c2.d.e[0].p1 == 2:
	}
}
