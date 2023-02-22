package linter

import (
	"testing"
)

func TestGoVersionParse(t *testing.T) {
	tests := []struct {
		s     string
		major int
		minor int
	}{
		{"", 0, 0},
		{"1.5", 1, 5},
		{"1.10", 1, 10},
		{"2.0", 2, 0},
		{"2.1", 2, 1},
	}

	runTest := func(x string, wantMajor, wantMinor int) {
		have, err := ParseGoVersion(x)
		if err != nil {
			t.Fatalf("parse %q: %v", x, err)
		}
		if have.Major != wantMajor {
			t.Errorf("parseGoVersion(%s); major: want %d, have %d", x, wantMajor, have.Major)
		}
		if have.Minor != wantMinor {
			t.Errorf("parseGoVersion(%s); minor: want %d, have %d", x, wantMinor, have.Minor)
		}
	}

	for _, test := range tests {
		runTest(test.s, test.major, test.minor)
		runTest("go"+test.s, test.major, test.minor)
	}
}

func TestGoVersionCompare(t *testing.T) {
	tests := []struct {
		version string
		other   string
		want    bool
	}{
		{"", "1.5", true},
		{"", "2.0", true},

		{"1.0", "1.0", true},
		{"1.0", "1.1", false},
		{"1.0", "2.0", false},

		{"1.16", "1.15", true},
		{"1.16", "1.16", true},
		{"1.16", "1.17", false},
		{"1.16", "2.0", false},
		{"1.16", "2.1", false},

		{"2.0", "1.0", true},
		{"2.0", "1.15", true},
		{"2.0", "1.254", true},
		{"2.0", "2.0", true},
		{"2.0", "2.1", false},
	}

	parseGoVersion := func(s string) GoVersion {
		v, err := ParseGoVersion(s)
		if err != nil {
			t.Fatalf("parse %q: %v", s, err)
		}
		return v
	}

	runTest := func(x, y string, want bool) {
		have := parseGoVersion(x).GreaterOrEqual(parseGoVersion(y))
		if have != want {
			t.Errorf("%s >= %s: incorrect result (want %v)", x, y, want)
		}
	}

	for _, test := range tests {
		runTest(test.version, test.other, test.want)
		runTest(test.version, "go"+test.other, test.want)
		runTest("go"+test.version, test.other, test.want)
		runTest("go"+test.version, "go"+test.other, test.want)
	}
}
