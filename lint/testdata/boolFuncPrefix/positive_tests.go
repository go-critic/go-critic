package checker_test

type foo struct{}

/// consider to add Is/Has/Contains prefix to function name
func enabled() bool { return true }

/// consider to add Is/Has/Contains prefix to function name
func (f *foo) active() bool { return true }
