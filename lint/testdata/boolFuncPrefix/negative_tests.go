package checker_test

func IsEnabled() bool { return true }

func HasString(s string, ss ...string) bool { return true }

func (f *foo) ContainsFlag(flag int) bool { return true }

func (f *foo) isEnabled() bool { return true }

func has(s string, ss ...string) bool { return true }

func (f *foo) containsFlag(flag int) bool { return true }
