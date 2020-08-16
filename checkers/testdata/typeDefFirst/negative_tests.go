package checker_test

type rec struct{}

func (r rec) Method() {}

// Considering no type declaration in file it is also ok
func (r recv) Method2() {}
