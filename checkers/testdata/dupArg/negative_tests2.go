package checker_test

import (
	"github.com/go-critic/go-critic/checkers/testdata/_importable/strings"
)

func nonStdPackage() {
	strings.Contains("")

	var x string
	strings.Contains(x, x)
}
