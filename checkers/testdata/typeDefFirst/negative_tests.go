package checker_test

type rec struct{}

func (r rec) Method()         {}
func (r *rec) MethodWithRef() {}

// Considering no type declaration in file it is also ok
func (r recv) Method2() {}

func JustFunction() {}
