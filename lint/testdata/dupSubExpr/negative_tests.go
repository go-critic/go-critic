package checker_test

func floatBinOps() {
	var x float64

	_ = x == x &&
		x != x &&
		x <= x &&
		x >= x
	_ = x / x
	_ = x - x
}

func noBinOpDuplicates() {
	var p point
	var xs [2]int

	if p.x|p.y == 0 {
	}

	if xs[0]&xs[1] == 1 {
	}

	if xs[1] < xs[0] {
	}

	if xs[1] > xs[0] {
	}

	if p.y == p.x || 1 < 2 {
	}

	if (1 + p.x + 3) >= (1 + p.y + 3) {
	}
}

func uncheckedOps(x int) {
	_ = x + x
	_ = x * x
}
