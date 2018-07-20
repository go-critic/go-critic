package checker_test

type object struct {
	a struct {
		c struct {
			x int
			d struct {
				e [1]struct {
					p1, p2, p3, p4 int
				}
			}
		}

		c1 struct {
			y int
			d struct {
				e [1]struct {
					p1, p2, p3, p4 int
				}
			}
		}

		c2 struct {
			z int
			d struct {
				e [1]struct {
					p1, p2, p3, p4 int
				}
			}
		}
	}
}

func newObj() *object {
	return &object{}
}

func repeatedSubExpr() {
	o := newObj()

	/// o.a.c.d.e[0].p1 repeated multiple times, consider assigning it to local variable
	_ = o.a.c.d.e[0].p1 + o.a.c.d.e[0].p1 + o.a.c.d.e[0].p1

	/// o.a.c.d.e[0] repeated multiple times, consider assigning it to local variable
	_ = o.a.c.d.e[0].p1 + o.a.c.d.e[0].p2 + o.a.c.d.e[0].p3

	/// (o.a.c.x + o.a.c.x) repeated multiple times, consider assigning it to local variable
	_ = (o.a.c.x + o.a.c.x) +
		(o.a.c.x + o.a.c.x) +
		(o.a.c.x + o.a.c.x)
}

func repeatedCaseExpr() {
	o := newObj()

	/// o.a.c.d.e[0] repeated multiple times, consider assigning it to local variable
	switch {
	case o.a.c.d.e[0].p1 == 0:
	case o.a.c.d.e[0].p2 == 1:
	case o.a.c.d.e[0].p3 == 2:
	}
}
