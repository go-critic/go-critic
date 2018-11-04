package checker_test

import (
	"fmt"
	"math"

	_ "github.com/go-toolsmith/astfmt" // To reproduce #665

	"github.com/go-critic/go-critic/lint"
)

func noWarnings() {
	fmt.Printf("Hello PI=%v, Rule=%v", math.Pi, lint.Rule{})
}

func noShadowByParams(x string, y int) (a string, b int) { return }

type noShadow struct{}

func (ns noShadow) f() {}

var _ = 10

func blankParam(_ int) {
	_ = 10
	_, x := 1, 2
	_ = x
}
