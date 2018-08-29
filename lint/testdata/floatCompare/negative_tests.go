package checker_test

func foo1() float64 {
	return 9.0
}

func g1() bool {
	var a, b float64 = 10.0, 20.0
	return a < b
}

func g2() bool {
	var a, b int = 10, 20
	return a == b
}

func g3() {
	var a, b float64 = 10.0, 20.0

	x := []float64{4.0, 5.0, 9.0, 10.0}

	var p *float64 = &a

	_ = b == 0.5

	_ = 'a' == 'b'

	_ = a >= b

	_ = a == b

	_ = a != 0.6

	_ = b < a && a >= b

	_ = a+b != a+b

	_ = a+b != (a + b)

	_ = (a + b) != (a + b)

	_ = foo1() == a

	_ = x[0] == a

	_ = *p == a

	_ = a == a+a

	_ = a+a != a

	_ = (a + a) == a

	_ = foo1()+foo1() == foo1()

	_ = x[0] != x[0]

	_ = x[0]+x[0] == x[0]

	_ = *p+*p == *p

	_ = x[0]+(x[0]) == x[0]

	_ = a+b == a+b
}
