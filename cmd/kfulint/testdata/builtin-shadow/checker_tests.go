package checker_test

func assigningToBuiltinFunctions() {
	///Shadowing: append shadowing
	append := 1
	_ = append

	///Shadowing: cap shadowing
	cap := 1
	_ = cap

	///Shadowing: close shadowing
	close := 1
	_ = close

	///Shadowing: complex shadowing
	complex := 1
	_ = complex

	///Shadowing: copy shadowing
	copy := 1
	_ = copy

	///Shadowing: delete shadowing
	delete := 1
	_ = delete

	///Shadowing: imag shadowing
	imag := 1
	_ = imag

	///Shadowing: len shadowing
	len := 1
	_ = len

	///Shadowing: make shadowing
	make := 1
	_ = make

	///Shadowing: new shadowing
	new := 1
	_ = new

	///Shadowing: panic shadowing
	panic := 1
	_ = panic

	///Shadowing: print shadowing
	print := 1
	_ = print

	///Shadowing: println shadowing
	println := 1
	_ = println

	///Shadowing: real shadowing
	real := 1
	_ = real

	///Shadowing: recover shadowing
	recover := 1
	_ = recover
}
