package checker_test

func (f foo) unexported() int { return 1 }
func (f foo) Exported() int   { return 1 }

func getFoo() foo { return foo{} }
