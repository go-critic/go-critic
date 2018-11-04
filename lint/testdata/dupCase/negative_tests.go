package checker_test

func goodBoolSwitch(x, y int, v []int) {
	switch {
	case true:
	case false:
	}

	switch true {
	case true, false:
	}

	switch {
	case x > 1:
	case y > 1:
	case x > 2:
	case y > 2:
	}

	switch true {
	case x > v[0], x > v[1], x > v[2], x > v[3], x > v[4]:
	case x > v[10]:
	}
}

func goodIntSwitch(x, y int, v []int) {
	switch x + 1 {
	case y:
	case v[0]:
	case v[1], v[2], v[3], v[4]:
	}

	switch 10 {
	case x + y + v[0]:
	case x + y + v[1]:
	case x + y + v[3]:
	case x + y + v[2]:
	}
}

func goodStructSwitch() {
	type point struct{ x, y int }

	switch (point{}) {
	case point{1, 2}:
	case point{2, 1}:
	case point{3, 3}:
	case point{0, 0}, point{10, 20}, point{20, 10}:
	default:
	}
}
