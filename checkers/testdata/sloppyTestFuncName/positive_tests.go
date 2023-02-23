package checker_test

import "testing"

/*! function TesstXXX should be of form TestXXX(t *testing.T) */
func TesstXXX(t *testing.T) {
	testCheck(t)
}

/*! function BenchXXX should be of form BenchmarkXXX(b *testing.B) */
func BenchXXX(b *testing.B) {
	benchCheck(b)
}

/*! function testCheck looks like a test helper, consider to change 1st param to 'tb testing.TB' */
func testCheck(t *testing.T) {
	t.Helper()

	if 1 == 0 {
		t.Fatal("really?")
	}
}

/*! function benchCheck looks like a benchmark helper, consider to change 1st param to 'tb testing.TB' */
func benchCheck(b *testing.B) {
	b.Helper()

	if 1 == 0 {
		b.Fatal("really?")
	}
}
