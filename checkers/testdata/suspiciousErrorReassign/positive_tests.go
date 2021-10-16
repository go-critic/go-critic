package checker_test

import (
	"fmt"
	"io"

	"github.com/go-critic/go-critic/checkers/testdata/suspiciousErrorReassign/pkg"
)

func _() {
	/*! suspicious reassigment of error from another package */
	pkg.ErrOur = nil

	/*! suspicious reassigment of error from another package */
	pkg.ErrOur = fmt.Errorf("your error is: %w", io.EOF)
}
