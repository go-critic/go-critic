package checker_test

import (
	"fmt"
)

var foo = "bar"

func fmtFormatting() {
	fmt.Errorf("%s", foo)
	fmt.Errorf("123")
	fmt.Errorf("%s", fooError())
}

func fooError() string { return "foo" }
