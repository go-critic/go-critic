package checker_test

import (
	/// unnecessary rename of import package
	fmt "fmt"

	/// unnecessary rename of import package
	lint "github.com/go-critic/go-critic/lint"

	/// unnecessary rename of import package
	dummypkg "github.com/go-critic/go-critic/lint/internal/dummy"
)

func warnings() {
	dummypkg.Dummy()

	fmt.Printf("Hello Rule=%v\n", lint.Rule{})
}
