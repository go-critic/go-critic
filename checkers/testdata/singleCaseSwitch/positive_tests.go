package checker_test

func intValue(x interface{}) int {
	/*! should rewrite switch statement to if statement */
	switch x := x.(type) {
	case int:
		return x
	}

	/*! should rewrite switch statement to if statement */
	switch x := x.(type) {
	case int:
		/*! should rewrite switch statement to if statement */
		switch x := interface{}(x).(type) {
		case int:
			println(x)
		}
	}

	return 0
}

func switchDefault(x interface{}) {
	/*! found switch with default case only */
	switch x.(type) {
	default:
	}

	/*! found switch with default case only */
	switch {
	default:
	}

	/*! found switch with default case only */
	switch {
	default:
		println(x)
		for {
			break
		}
	}

	/*! found switch with default case only */
	switch x := x.(type) {
	default:
		/*! should rewrite switch statement to if statement */
		switch x {
		case 1:
		}
	}

	/*! found switch with default case only */
	switch x := x.(type) {
	default:
		/*! should rewrite switch statement to if statement */
		switch d := x.(type) {
		case int:
			println(d)
		}
	}

	/*! found switch with default case only */
	switch x := x.(type) {
	default:
		switch x {
		case 1:
		case 2:
		}
	}
}

func switchWithOneCase(x int) {
	/*! should rewrite switch statement to if statement */
	switch x {
	case 1:
	}

	/*! should rewrite switch statement to if statement */
	switch {
	case true:
	}

	/*! should rewrite switch statement to if statement */
	switch x == 1 {
	case true:
	}

	/*! should rewrite switch statement to if statement */
	switch interface{}(x).(type) {
	case int32:
	}

	/*! should rewrite switch statement to if statement */
	switch x {
	case 1:
		/*! should rewrite switch statement to if statement */
		switch x {
		case 2:
		}
	}

	/*! should rewrite switch statement to if statement */
	switch x {
	case 1:
		switch x {
		case 2:
		case 3:
		}
	}
}

func badCaseWithBreak(x, y int) {
	/*! should rewrite switch statement to if statement */
	switch x {
	case 0:
		println(x)
		for {
			break
		}
	}

	/*! should rewrite switch statement to if statement */
	switch x {
	case 0:
		println(x)
		switch y {
		case 2:
			break
		}
	}

	/*! should rewrite switch statement to if statement */
	switch x {
	case 0:
		println(x)
		switch y {
		case 2:
			break
		case 1:
			break
		}
	}
}

func badCaseWithRange(a []string) {
	/*! should rewrite switch statement to if statement */
	switch true {
	case true:
		for range a {
		}
	}

	/*! should rewrite switch statement to if statement */
	switch x := interface{}(a).(type) {
	case uint:
		for range a {
			println(x)
		}
	}

	/*! should rewrite switch statement to if statement */
	switch a[0] {
	case "2":
		for range a {
			println(a)
		}
	}

	for i := range a {
		/*! should rewrite switch statement to if statement */
		switch i {
		case 1:
		}
	}

	for i := range a {
		/*! found switch with default case only */
		switch i {
		default:
		}
	}
}

func badCaseWithSelect(ch chan string) {
	/*! should rewrite switch statement to if statement */
	switch {
	case true:
		select {
		case <-ch:
		}
	}

	/*! should rewrite switch statement to if statement */
	switch x := interface{}(0).(type) {
	case string:
		select {
		case ch <- x:
		}
	}

	/*! should rewrite switch statement to if statement */
	switch <-ch {
	case "1":
		select {
		case ch <- "1":
		}
	}

	/*! found switch with default case only */
	switch {
	default:
		select {
		default:
		}
	}

	select {
	case <-ch:
		/*! found switch with default case only */
		switch {
		default:
		}
	default:
		/*! should rewrite switch statement to if statement */
		switch {
		case true:
		}
	}

	select {
	default:
		/*! should rewrite switch statement to if statement */
		switch xx := interface{}(ch).(type) {
		case int:
			println(xx)
		}
	}

	select {
	case ch <- "2":
		/*! should rewrite switch statement to if statement */
		switch <-ch {
		case "1":
		}
	}
}
