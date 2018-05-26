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

	"github.com/PieselBois/kfulint/lint"
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

type ruleTest struct {
	name string
}

var ruleList []ruleTest

func init() {
	for _, rule := range lint.RuleList() {
		ruleList = append(ruleList, ruleTest{
			name: rule.Name(),
		})
	}
}

func runChecker(name, pkgPath string) (output []byte, err error) {
	return exec.Command("./"+binary, "-enable", name, pkgPath).CombinedOutput()
}

func TestSanity(t *testing.T) {
	saneRules := ruleList[:0]
	for _, rule := range ruleList {
		pkgPath := linterCmdPath + "testdata/_sanity"
		output, err := runChecker(rule.name, pkgPath)
		if err != nil {
			t.Errorf("%s failed sanity checks: %v:\n%s",
				rule.name, err, output)
			continue
		}
		saneRules = append(saneRules, rule)
	}
	ruleList = saneRules
}

func TestOutput(t *testing.T) {
	for _, rule := range ruleList {
		t.Run(rule.name, func(t *testing.T) {
			pkgPath := linterCmdPath + "testdata/" + rule.name
			testFilename := filepath.Join(
				"testdata", rule.name, "checker_tests.go")
			f := parseTestFile(t, testFilename)

			// Running the linter.
			output, err := runChecker(rule.name, pkgPath)
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
				var lineString, ruleName, text string
				unpackSubmatches(l, warningRE, &lineString, &ruleName, &text)
				line, err := strconv.Atoi(lineString)
				if err != nil {
					t.Errorf("%s: invalid line number in %s",
						testFilename, lineString)
				}
				if ruleName != rule.name {
					t.Errorf("%s: unexpected checker name: %s",
						testFilename, ruleName)
					continue
				}
				if w := f.Find(line, text); w != nil {
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
	// Example: "/// can replace s[:] with s".
	warningDirectiveRE = regexp.MustCompile(`^\s*/// (.*)`)
)

// warning is a decoded warning directive that is used to match actual
// linter warnings against expected results.
type warning struct {
	matched bool
	text    string
}

func (w warning) String() string {
	return "/// " + w.text
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
		unpackSubmatches(l, warningDirectiveRE, &m.text)
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
func (f *testFile) Find(line int, text string) *warning {
	for _, y := range f.warnings[line] {
		if text == y.text {
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
