package checker_test

type pair struct {
	first  int
	second int
}

func fieldSwap(p *pair) {
	/*! can re-write as `p.first, p.second = p.second, p.first` */
	tmp := p.first
	p.first = p.second
	p.second = tmp
}

func varSwap(x, y int) {
	/*! can re-write as `x, y = y, x` */
	tmp := x
	x = y
	y = tmp
}

func pointersSwap1(x, y *int) {
	/*! can re-write as `*x, *y = *y, *x` */
	tmp := *x
	*x = *y
	*y = tmp
}

func pointersSwap2(x, y *int) {
	/*! can re-write as `*y, *x = *x, *y` */
	tmp := *y
	*y = *x
	*x = tmp
}
