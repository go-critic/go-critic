package checker_test

type t struct {
	a struct {
		c struct {
			d struct {
				e struct {
					p1, p2, p3, p4 int
				}
			}
		}

		c1 struct {
			d struct {
				e struct {
					p1, p2, p3, p4 int
				}
			}
		}

		c2 struct {
			d struct {
				e struct {
					p1, p2, p3, p4 int
				}
			}
		}
	}
}

func sample() {
	a := 1
	b := t{}
	///: Expression chain b.a.c.d.e repeated multiple times consider assigning it to local variable
	switch a {
	case b.a.c.d.e.p1:
	case b.a.c.d.e.p2:
	case b.a.c.d.e.p3:
	case b.a.c.d.e.p4:
	}

	switch a {
	case b.a.c.d.e.p1:
	case b.a.c1.d.e.p1:
	case b.a.c2.d.e.p1:
	}
}
