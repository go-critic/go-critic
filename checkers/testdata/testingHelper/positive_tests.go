package checker_test

import (
	"testing"
)

/*! consider to make *testing.T a first parameter */
func bad1(a int, t *testing.T) {}

/*! consider to make *testing.T a first parameter */
func bad2(a int, s string, t *testing.T, _ float32) {}

/*! consider to call t.Helper() a first statement */
func bad3(t *testing.T) {
	type in = int

	t.Helper()
}

/*! consider to call t.Helper() a first statement */
func bad4(t *testing.T) {
	t.Cleanup()

	t.Helper()
}
