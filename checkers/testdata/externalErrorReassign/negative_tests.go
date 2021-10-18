package checker_test

import (
	"github.com/go-critic/go-critic/checkers/testdata/_importable/examplepkg"
)

func _() {
	examplepkg.ReassignFoo(nil)
}
