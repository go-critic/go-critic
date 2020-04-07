package checker_test

import "fmt"

func f(...interface{}) int { return 1 }

const ten = 10

func positiveTests() {
	/*! can rewrite as `defer f()` */
	defer func() { f() }()

	/*! can rewrite as `defer f(1)` */
	defer func() { f(1) }()

	/*! can rewrite as `defer f(ten, ten+1)` */
	defer func() { f(ten, ten+1) }()

	/*! can rewrite as `defer fmt.Println("hello")` */
	defer func() { fmt.Println("hello") }()
}
