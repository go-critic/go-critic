package checker_test

func positive () {
	a := 1
	/// block doesn't have assignment statements, can be simply deleted
	{
		print(a)
	}
}
