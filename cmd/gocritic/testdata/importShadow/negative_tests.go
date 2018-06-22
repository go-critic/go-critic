package checker_test

import (
	"fmt"
	"math"

	"github.com/go-critic/go-critic/lint"
)

func noWarnings() {
	fmt.Printf("Hello PI=%v, Rule=%v", math.Pi, lint.Rule{})
}

func noShadowByParams(x string, y int) (a string, b int) { return }

type noShadow struct{}

func (ns noShadow) f() {}
