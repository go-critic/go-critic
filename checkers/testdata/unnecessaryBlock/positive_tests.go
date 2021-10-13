package checker_test

func positive() {
	a := 1
	/*! block doesn't have definitions, can be simply deleted */
	{
		print(a)
		/*! block doesn't have definitions, can be simply deleted */
		{
			println("2")
		}
	}

	/*! block doesn't have definitions, can be simply deleted */
	{
		a = 10 // Not a definition
		print(a)
	}

	/*! block doesn't have definitions, can be simply deleted */
	{
		type ()
		println("empty type decl (0 specs)")
		type ()
	}

	/*! block doesn't have definitions, can be simply deleted */
	{
		var ()
		println("empty var decl (0 specs)")
		var ()
	}

	/*! block doesn't have definitions, can be simply deleted */
	{
		const ()
		println("empty const decl (0 specs)")
		const ()
	}

	switch a {
	case 1:
		/*! case statement doesn't require a block statement */
		{
			x := 10
			println(x)
		}
	}

	var ch chan string
	select {
	case ch <- "1":
		/*! case statement doesn't require a block statement */
		{
			x := 10
			println(x)
		}
	}
}
