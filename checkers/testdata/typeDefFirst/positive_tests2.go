package checker_test

func (r *reciv) Method() {}

/*! definition of type 'reciv' should appear before its methods */
type reciv struct{ x, y int }
