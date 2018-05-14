package checker_test

type t struct {
	a struct {
		b struct {
			c struct {
				d struct {
					e struct {
						p1, p2, p3, p4 int
					}
				}
			}
		}
	}
}

func sample() {
	a := 1
	b := t{}
	///: Expression chain b.a.b.c.d.e repeated multiple times consider assigning it to local variable
	switch a {
	case b.a.b.c.d.e.p1:
	case b.a.b.c.d.e.p2:
	case b.a.b.c.d.e.p3:
	case b.a.b.c.d.e.p4:
	}
}
