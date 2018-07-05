package checker_test

func suggestSwitch() {
	cond1 := true
	cond2 := true
	cond3 := true

	/// should rewrite if-else to switch statement
	if cond1 {
	} else if cond2 {
	} else {
	}

	/// should rewrite if-else to switch statement
	if cond1 {
	} else if cond2 {
	} else if cond3 {
	}

	/// should rewrite if-else to switch statement
	if cond1 {
	} else if cond2 {
	} else if cond3 {
	} else {
	}

	/// should rewrite if-else to switch statement
	if cond1 {
	} else if cond2 {
		if cond3 {
		}

		/// should rewrite if-else to switch statement
		if cond1 {
		} else if cond2 {
		} else if cond3 {
		} else {
		}
	} else {
		/// should rewrite if-else to switch statement
		if cond1 {
		} else if cond2 {
		} else {
		}
	}
}

func describeInt(x int) string {
	/// should rewrite if-else to switch statement
	if x == 0 {
		return "zero"
	} else if x < 0 {
		return "negative"
	} else {
		return "positive"
	}
}
