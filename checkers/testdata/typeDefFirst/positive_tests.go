package checker_test

func (r recv) MethodBefore() {}

/*! definition of type 'recv' should appear before its methods */
type recv struct{}
