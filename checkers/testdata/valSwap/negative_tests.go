package checker_test

func fieldSwapOK(p *pair) {
	p.first, p.second = p.second, p.first
}

func varSwapOK(x, y int) {
	x, y = y, x
}

func pointersSwap1OK(x, y *int) {
	*x, *y = *y, *x
}

func pointersSwap2OK(x, y *int) {
	*y, *x = *x, *y
}

func notSwap(x, y int) {
	tmp := x
	y = x // logic error here, probably...
	y = tmp
}
