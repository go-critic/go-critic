package checker_test

import (
	"github.com/go-critic/go-critic/checkers/testdata/suspiciousErrorReassign/pkg"
)

func _() {
	pkg.ReassignError(nil)
}
