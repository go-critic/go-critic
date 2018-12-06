package checker_test

import (
	"fmt"
	"math"

	mymath1 "math"
	mymath2 "math"

	"github.com/go-lintpack/lintpack"
)

var _ = mymath1.E
var _ = mymath2.E

func shadowImportedPackages() {
	fmt.Printf("Hello PI=%v, Info=%v", math.Pi, lintpack.CheckerInfo{})

	/*! shadow of imported package 'math' */
	math := "some math"
	/*! shadow of imported from 'github.com/go-lintpack/lintpack' package 'lintpack' */
	lintpack := "some lint"

	fmt.Printf("Hello math=%v, lint=%v", math, lint)
}

func genDeclShadow() {
	/*! shadow of imported package 'math' */
	const math = 1
	var (
		/*! shadow of imported package 'fmt' */
		fmt = 2
		/*! shadow of imported from 'github.com/go-lintpack/lintpack' package 'lintpack' */
		lintpack = 3
	)
	_, _ = fmt, lintpack
}

/*! shadow of imported package 'math' */
/*! shadow of imported package 'fmt' */
func shadowedByParam1(math string, fmt int) {}

/*! shadow of imported package 'math' */
/*! shadow of imported package 'fmt' */
func shadowedByParam2() (math string, fmt int) { return }

type shadower struct{}

/*! shadow of imported package 'fmt' */
func (fmt shadower) f() {}

/*! shadow of imported package 'mymath1' */
func renamedImportShadow1(mymath1 int) {}

/*! shadow of imported package 'mymath2' */
func renamedImportShadow2(mymath2 int) {}
