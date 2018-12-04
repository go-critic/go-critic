package bar_test

import "testing"

func TestExternal(t *testing.T) {
	var object *struct{ x int }
	_ = (*object).x
}
