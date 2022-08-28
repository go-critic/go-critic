package checker_test

import (
	"fmt"
	"reflect"
)

func _() {
	{
		var foo withStringer
		_ = foo.String()
	}

	{
		var err error
		_ = err.Error()
	}

	{
		var s string
		_ = s

		_ = "x"
	}

	{
		var rv reflect.Value
		_ = fmt.Sprint(rv)
		_ = fmt.Sprintf("%s", rv)
		_ = fmt.Sprintf("%v", rv)
	}
}
