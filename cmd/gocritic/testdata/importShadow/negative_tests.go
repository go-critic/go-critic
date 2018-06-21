package checker_test

import (
	"fmt"
	"math"

	"github.com/go-critic/go-critic/lint"
)

func noWarnings() {
	fmt.Printf("Hello PI=%v, Rule=%v", math.Pi, lint.Rule{})
}
