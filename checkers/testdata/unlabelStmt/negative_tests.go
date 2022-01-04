package checker_test

func usedLabels(v interface{}) {
label2:
	for {
		for {
			break label2
		}
		break label2
	}

label3:
	switch {
	case true:
		for {
			{
				break label3
			}
		}
	case false:
		if true {
			break label3
		}
	}

label4:
	switch v.(type) {
	case int:
		select {
		default:
			break label4
		}
	default:
		break label4
	}
}

func unlabeled(v interface{}) {
	for {
		break
	}

	for {
		for {
			break
		}
		break
	}

	switch {
	case true:
		for {
			{
				break
			}
		}
	case false:
		if true {
			break
		}
	}

	switch v.(type) {
	case int:
		select {
		default:
			break
		}
	default:
		break
	}
}

func withGoto() {
	_ = func() {
	OuterFor:
		for {
			break OuterFor
		}
		goto OuterFor
	}

	_ = func() {
	OuterRange:
		for range []int{} {
			continue OuterRange
		}
		goto OuterRange
	}

	_ = func() {
	OuterSwitch:
		switch {
		default:
			break OuterSwitch
		}
		goto OuterSwitch
	}

	_ = func() {
	OuterTypeSwitch:
		switch interface{}(1).(type) {
		default:
			break OuterTypeSwitch
		}
		goto OuterTypeSwitch
	}
}

func forContinue(xs, ys []int) {
	for range xs {
		for range ys {
			continue
		}
	}

	for _, x := range xs {
		_ = x
		for i := 0; i < len(ys); i++ {
			continue
		}
	}

	for range xs {
		for range ys {
			break
		}
	}

	for _, x := range xs {
		_ = x
		for i := 0; i < len(ys); i++ {
			break
		}
	}

	for range xs {
		for range ys {
		}
	}

outer2:
	for _, x := range xs {
		for i := 0; i < len(ys); i++ {
			continue outer2
		}
		_ = x
	}
}

func twoLoopsWithSelect(c chan int) {
outer2:
	for {
		println("foo")
		for {
			select {
			case <-c:
				continue outer2
			}
		}
	}
}
