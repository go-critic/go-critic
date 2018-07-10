package checker_tests

func idealExpressions() {
	// No negations:
	_ = true
	_ = true
}

func negationOK() {
	_ = !true
	_ = !false
}

func otherBinOps() {
	_ = true != false
	_ = true && false
	_ = true || false
}
