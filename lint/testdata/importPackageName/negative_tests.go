package checker_test

import (
	"fmt"

	linter "github.com/go-critic/go-critic/lint"

	dummy "github.com/go-critic/go-critic/lint/internal/dummy"
)

func noWarnings() {
	dummy.Dummy()

	fmt.Printf("Hello Rule=%v", linter.Rule{})
}
