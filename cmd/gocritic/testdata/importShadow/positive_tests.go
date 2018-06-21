package checker_test

import (
	"fmt"
	"math"

	"github.com/go-critic/go-critic/lint"
)

func shadowImportedPackages() {
	fmt.Printf("Hello PI=%v, Rule=%v", math.Pi, lint.Rule{})

	/// shadow of imported package 'math'
	math := "some math"
	/// shadow of imported from 'github.com/go-critic/go-critic/lint' package 'lint'
	lint := "some lint"

	fmt.Printf("Hello math=%v, lint=%v", math, lint)
}
