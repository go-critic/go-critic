package main_test

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

const (
	// binary is test linter executable name.
	// Using "exe" suffix to make it work on Windows as well.
	binary = "testlint.exe"

	// linterCmdPath holds full path to linter main pkg path.
	linterCmdPath = "github.com/PieselBois/kfulint/cmd/kfulint/"
)

func TestMain(m *testing.M) {
	// Before all tests are executed, we need to build linter first.
	// After all tests are completed, linter executable is removed.

	buildLinter()
	// Run all tests in a normal way.
	exitStatus := m.Run()
	// Nothing bad will happen if someone already removed the binary.
	_ = os.Remove(binary)

	os.Exit(exitStatus)
}

var tests = []*struct {
	checker string
}{
	{"param-name"},
	{"type-guard"},
	{"parenthesis"},
	{"param-duplication"},
	{"underef"},
	{"elseif"},
	{"big-copy"},
	{"long-chain"},
}

func runChecker(name, pkgPath string) (output []byte, err error) {
	return exec.Command("./"+binary, "-enable", name, pkgPath).CombinedOutput()
}

func TestSanity(t *testing.T) {
	saneTests := tests[:0]
	for _, test := range tests {
		pkgPath := linterCmdPath + "testdata/_sanity"
		output, err := runChecker(test.checker, pkgPath)
		if err != nil {
			t.Errorf("%s failed sanity checks: %v:\n%s",
				test.checker, err, output)
			continue
		}
		saneTests = append(saneTests, test)
	}
	tests = saneTests
}

func TestOutput(t *testing.T) {
	for _, test := range tests {
		t.Run(test.checker, func(t *testing.T) {
			pkgPath := linterCmdPath + "testdata/" + test.checker
			testFilename := filepath.Join(
				"testdata", test.checker, "checker_tests.go")
			f := parseTestFile(t, testFilename)

			// Running the linter.
			output, err := runChecker(test.checker, pkgPath)
			if err != nil {
				t.Fatalf("run linter: %v: %s", err, output)
			}

			warningRE := regexp.MustCompile(`.*?:(\d+):\d+: ([a-zA-Z\-/]*): (.*)`)

			// Process linter output.
			var unexpectedLines []string
			for _, l := range strings.Split(string(output), "\n") {
				if len(l) == 0 { // Ignore empty lines
					continue
				}
				if !warningRE.MatchString(l) {
					// Something that doen't look like a warning.
					// Probably debug output or checker runtime error.
					unexpectedLines = append(unexpectedLines, l)
					continue
				}
				var lineString, tag, text string
				unpackSubmatches(l, warningRE, &lineString, &tag, &text)
				line, err := strconv.Atoi(lineString)
				if err != nil {
					t.Errorf("%s: invalid line number in %s",
						testFilename, lineString)
				}
				checker, kind := splitWarningTag(tag)
				if checker != test.checker {
					t.Errorf("%s: unexpected checker name: %s",
						testFilename, checker)
					continue
				}
				if w := f.Find(line, kind, text); w != nil {
					if w.matched {
						t.Errorf("%s:%d: multiple matches for %s",
							testFilename, line, w)
					}
					w.matched = true
				} else {
					unexpectedLines = append(unexpectedLines, l)
				}
			}

			// Check if there are unmatched warnings and/or unexpected
			// lines in the linter output.
			for line := range f.warnings {
				for _, w := range f.warnings[line] {
					if w.matched {
						continue
					}
					t.Errorf("%s:%d: unmatched %s", testFilename, line, w)
				}
			}
			for _, l := range unexpectedLines {
				t.Errorf("unexpected line in output: %s", l)
			}
		})
	}
}

// parseTestFile decodes single end-to-end test file.
// Errors reported through t.
func parseTestFile(t *testing.T, filename string) *testFile {
	p := testFileParser{
		t:        t,
		filename: filename,
	}
	return p.Parse()
}

var (
	// warningDirectiveRE describes pattern used to match "///" directives inside
	// end-to-end test files.
	//
	// Directive line contain only special comment itself, no other
	// syntax elements (whitespace is permitted).
	//
	// matcher = "///" [kind] ":" text
	// kind = \w*
	// text = .*
	//
	// Example: "///Label: remove unused label"
	// Example: "///: can replace s[:] with s"
	warningDirectiveRE = regexp.MustCompile(`^\s*///(\w*?): (.*)`)
)

// warning is a decoded warning directive that is used to match actual
// linter warnings against expected results.
type warning struct {
	matched bool
	kind    string
	text    string
}

func (w warning) String() string {
	return "///" + w.kind + ": " + w.text
}

type testFileParser struct {
	t *testing.T

	filename string
}

func (p *testFileParser) Parse() *testFile {
	f := &testFile{warnings: make(map[int][]*warning)}
	lines := p.readFileLines()
	p.collectWarnings(lines, f)
	return f
}

// collectWarnings scans lines for "expected warning" directives and
// fills f with them.
func (p *testFileParser) collectWarnings(lines []string, f *testFile) {
	var pending []*warning
	for i, l := range lines {
		if !warningDirectiveRE.MatchString(l) {
			if len(pending) != 0 {
				line := i + 1
				f.warnings[line] = append([]*warning{}, pending...)
				pending = pending[:0]
			}
			continue
		}
		var m warning
		unpackSubmatches(l, warningDirectiveRE, &m.kind, &m.text)
		pending = append(pending, &m)
	}
}

// readFileLines returns associated file contents as a slice of lines.
func (p *testFileParser) readFileLines() []string {
	testData, err := ioutil.ReadFile(p.filename)
	if err != nil {
		p.t.Fatalf("can't find checker tests: %v", err)
	}
	return strings.Split(string(testData), "\n")
}

// testFile is decoded testdata file.
type testFile struct {
	warnings map[int][]*warning
}

// Find seeks for matching warning.
func (f *testFile) Find(line int, kind, text string) *warning {
	for _, y := range f.warnings[line] {
		if kind == y.kind && text == y.text {
			return y
		}
	}
	return nil
}

// unpackSubmatches binds re.FindStringSubmatch(s) results to dst strings.
func unpackSubmatches(s string, re *regexp.Regexp, dst ...*string) {
	submatches := re.FindStringSubmatch(s)
	// Skip [0] which is a "whole match".
	for i, submatch := range submatches[1:] {
		*dst[i] = submatch
	}
}

// splitWarningTag returns warning tag components, namely: checker name and warning kind.
func splitWarningTag(tag string) (checkerName, warningKind string) {
	parts := strings.Split(tag, "/")
	checkerName = parts[0]
	if len(parts) > 1 {
		warningKind = parts[1]
	}
	return checkerName, warningKind
}

func buildLinter() {
	// TODO(quasilyte): check that this actually works on windows.
	goBin := "go"
	if runtime.GOOS == "windows" {
		goBin += ".exe"
	}
	output, err := exec.Command(goBin, "build", "-o", binary).CombinedOutput()
	if err != nil {
		log.Fatalf("build linter: %v: %s", err, output)
	}
}
