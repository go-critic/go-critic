package checker_test

func intValue1(x interface{}) int {
	switch x := x.(type) {
	case int:
		return x
	case uint:
		return 1
	}
	return 0
}

func switchWithOneCaseAndDefault(x int) {
	switch x {
	default:
	case 1:
	}
}

func switchWithTwoCases(x int) {
	switch x {
	case 1:
	case 2:
	}
}

func caseWithTwoValues(x int) {
	switch x {
	case 1, 2:
	}

	switch xx := interface{}(x).(type) {
	case int, uint:
		println(xx)
	}

	switch {
	case false, true:
	}

	switch 1 {
	case 2, 3:
	}
}

func caseWithBreak(x interface{}) {
	switch x.(type) {
	case int:
		println(x)
		break
	}

	switch x.(int) {
	case 0:
		println(x)
		break
	}

	switch xx := x.(type) {
	default:
		println(xx)
		break
	}

	for {
		switch x.(int) {
		case 0:
			println(x)
			break
		}
	}
}

func caseWithSelect(ch chan string) {
	switch <-ch {
	case "1", "2":
		select {
		case ch <- "1":
		}
	}

	switch <-ch {
	case "1":
		select {
		case ch <- "1":
		}
	default:
		select {
		default:
		}
	}

	switch x := interface{}(<-ch).(type) {
	case int, uint:
		select {
		case ch <- "1":
			println(x)
		}
	}

	select {
	case <-ch:
		switch <-ch {
		default:
		case "1":
		}
	}

	select {
	default:
		switch <-ch {
		case "1", "2":
		}
	}
}

func caseWithRange(a []string) {
	switch true {
	case true:
		for range a {
		}
	default:
	}

	switch x := interface{}(a).(type) {
	case uint, int:
		for range a {
			println(x)
		}
	}

	switch a[0] {
	case "2", "3":
		for range a {
			println(a)
		}
	}

	for i := range a {
		switch i {
		case 1, 2:
		}
	}

	for i := range a {
		switch i {
		default:
			break
		}
	}
}
