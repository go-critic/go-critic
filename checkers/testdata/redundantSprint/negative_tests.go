package checker_test

func _() {
	{
		var foo withStringer
		_ = foo.String()
	}

	{
		var err error
		_ = err.Error()
	}

	{
		var s string
		_ = s

		_ = "x"
	}
}
