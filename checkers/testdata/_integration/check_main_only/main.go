package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
)

func appendAssign() []string {
	var xs, ys []string
	xs = append(ys, "x")
	return xs
}

func appendCombine(xs []int) {
	xs = append(xs, 1)
	xs = append(xs, 2)
}

func assignOp(x int) {
	x = x + 2
}

func boolExprSimplify(x, y int) bool {
	return !(x == y+1)
}

func builtinShadow(new int) {}

func captLocal(THIS int) {}

func caseOrder(x interface{}) {
	switch x.(type) {
	case interface{}:
	case int:
	}
}

func commentedOutCode() {
	// x := foo.bar(arg, arg, "123")
}

func defaultCaseOrder(x int) string {
	switch x {
	case 10:
		return "ten"
	default:
		return "?"
	case 1:
		return "one"
	}
}

// DeprecatedComment example of deprecated comment done wrong.
//
// Deprecated. Use abc instead.
func DeprecatedComment() {}

// DocStub ...
func DocStub() {}

func dupArg(xs []int) {
	copy(xs, xs)
}

func dupBranchBody(x int) int {
	if x == 0 {
		x++
	} else {
		x++
	}
	return x
}

func dupCase(x int) {
	switch {
	case x == 0:
	case x == 0:
	}
}

func dupSubExpr(x int) bool {
	return x*x < x*x
}

func elseif(cond1, cond2 bool) {
	if cond2 {
	} else {
		if cond2 {
		}
	}
}

func emptyFallthrough(x int) string {
	switch x {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		return "!"
	default:
		return "?"
	}
}

func flagDeref() {
	_ = *flag.String("str", "", "usage")
}

func hugeParam(x [10000]int) {}

func ifElseChain(cond1, cond2, cond3 bool) {
	if cond1 {
	} else if cond2 {
	} else if cond3 {
	}
}

func importShadow(flag int) {}

func indexAlloc(s []byte, sub string) {
	_ = strings.Index(string(s), sub)
}

func methodExprCall(p point) {
	_ = point.String(p)
}

func nilValReturn(x *int) *int {
	if x == nil {
		return x
	}
	return new(int)
}

func paramTypeCombine(x int, y int) {}

func rangeExprCopy(xs [1000]int) {
	for _, x := range xs {
		_ = x
	}
}

func rangeValCopy(xs [][1000]int) {
	for _, x := range xs {
		_ = x
	}
}

func regexpMust() {
	re, _ := regexp.Compile(`this`)
	_ = re
}

func singleCaseSwitch(x int) string {
	switch x {
	case 0:
		return "0"
	}
	return fmt.Sprint(x)
}

func sloppyLen(xs []int) bool {
	return len(xs) < 0
}

func sloppyReassign() error {
	var err error
	if err = (point{}); err != nil {
		return err
	}
	return nil
}

func switchTrue(x int) string {
	switch true {
	case x > 10:
		return "more than 10"
	default:
		return "?"
	}
}

func typeSwitchVar(x interface{}) int {
	switch x.(type) {
	case int:
		return x.(int) * 2
	case float32:
		return int(x.(float32)) * 2
	default:
		return -1
	}
}

func typeUnparen() {
	var _ (func())
}

func underef(xs *[10]int) int {
	return (*xs)[2]
}

func unlabelStmt() {
loop:
	for {
		continue loop
	}
}

func unlambda(func(x int) int) {
	// Don't use local func here, see #888.
	unlambda(func(x int) int { return add1(x) })
}

func unslice(xs []int) []int { return xs[:] }

func wrapperFunc(wg *sync.WaitGroup) {
	wg.Add(-1)
}

type point struct{ x, y int }

func (p point) Error() string { return "just a point!" }

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func offBy1(xs []int) int {
	return xs[len(xs)]
}

func flagName() {
	_ = flag.String(" foo ", "1", "")
}

func commentFormatting() {
	//"123"
}

func badCond(x int) {
	_ = x < 100 && x > 200
}

func weakCond(xs []int) {
	_ = xs == nil || xs[0] == 0
}

func exitAfterDefer() {
	defer func() {}()
	log.Fatal(123)
}

func typeAssertChain() {
	var x interface{}
	if v, ok := x.(int8); ok {
		_ = v
	} else if v, ok := x.(int16); ok {
		_ = v
	}
}

func argOrder(s string) {
	_ = strings.HasPrefix("$", s)
}

func newDeref() {
	_ = *new(string)
}

func badCall(s string) {
	_ = strings.Replace(s, "-", "=", 0)
}

// No test functions below this line, please.

func main() {
}

func add1(x int) int { return x + 1 }
