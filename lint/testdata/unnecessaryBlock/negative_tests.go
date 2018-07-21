package checker_test

func negative() {
	{
		a := 1
		print(a)
	}

	{
		var a = 10
		print(a)
	}

	{
		const a = 10
		print(a)
	}

	switch {
	case false:
		println("block inside case clause without defs is OK")
	default:
		println("block inside default case clause without defs is OK")
	}

	if false {
		println("block inside if stmt without defs is OK")
	}

	for {
		println("block inside for stmt without defs is OK")
		break
	}

	for {
		for {
			println("nested blocks are also OK")
			break
		}
		break
	}
}

func withTypeDef() {
	{
		type foo int
		println(foo(1))
	}
	{
		type foo float64
		println(foo(1.1))
	}
}
