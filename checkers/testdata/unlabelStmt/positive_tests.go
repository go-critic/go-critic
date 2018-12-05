package checker_test

func redundantLabels(v interface{}, xs []int) {
	/*! label label1 is redundant */
label1:
	for false {
		break label1
	}

	/*! label label2 is redundant */
label2:
	for {
		for range xs {
			break
		}
		break label2
	}

	/*! label label3 is redundant */
label3:
	switch {
	case true:
		{
			break label3
		}
	case false:
		if true {
			break label3
		}
	}

	/*! label label4 is redundant */
label4:
	switch v.(type) {
	case int:
		select {
		default:
			break
		}
	default:
		break label4
	}
}

func forContinueOuter(xs, ys []int) {
outer1:
	for range xs {
		for range ys {
			/*! change `continue outer1` to `break` */
			continue outer1
		}
	}

outer2:
	for _, x := range xs {
		_ = x
		for i := 0; i < len(ys); i++ {
			/*! change `continue outer2` to `break` */
			continue outer2
		}
	}
}
