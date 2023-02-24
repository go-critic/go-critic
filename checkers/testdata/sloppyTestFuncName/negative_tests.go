package checker_test

import "testing"

func TestXXX(t *testing.T) {
	testHelper(t, 1)
}

func BenchmarkXXX(b *testing.B) {
	benchHelper(b, 1)
}

func testHelper(t *testing.T, a int) {
	t.Helper()

	if a == 0 {
		t.Fatal("really?")
	}
}

func benchHelper(b *testing.B, a int) {
	b.Helper()

	if a == 0 {
		b.Fatal("really?")
	}
}

func skiptestXXX(t *testing.T) {
	// blabla
}

func skipbenchXXX(b *testing.B) {
	// blabla
}
