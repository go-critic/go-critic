package checker_test

import (
	"fmt"
	"math"

	"github.com/go-critic/go-critic/linter"

	_ "github.com/go-toolsmith/astfmt" // To reproduce #665
)

func noWarnings() {
	fmt.Printf("Hello PI=%v, Info=%v", math.Pi, linter.CheckerInfo{})
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
