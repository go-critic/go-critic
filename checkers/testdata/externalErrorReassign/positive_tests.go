package checker_test

import (
	"fmt"
	"io"

	"github.com/go-critic/go-critic/checkers/testdata/_importable/examplepkg"
)

func _() {
	/*! suspicious reassigment of error from another package */
	examplepkg.FooError = nil

	/*! suspicious reassigment of error from another package */
	examplepkg.FooError = fmt.Errorf("your error is: %w", io.EOF)
}
