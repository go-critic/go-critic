package checker_test

import (
	"testing"
)

/*! FOO */
func t1(a int, t *testing.T) {}

/*! FOO */
func t2(a int, s string, t *testing.T, _ float32) {}

/*! FOO */
func t3(t *testing.T) {
	type in = int

	t.Helper()
}

/*! FOO */
func t4(t *testing.T) {
	t.Cleanup()

	t.Helper()
}

/*! FOO */
func t5(t *testing.T) {}

/*! FOO */
func t6(t *testing.T) {
	// just a comment
	t.Helper()
}
