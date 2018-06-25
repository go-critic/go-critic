package lint

import (
	"go/ast"
	"go/parser"
	"go/types"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/tools/go/loader"
)

var (
	testdataPkgPath    = "github.com/go-critic/go-critic/lint/testdata/"
	sizes              = types.SizesFor("gc", runtime.GOARCH)
	warningRE          = regexp.MustCompile(`.*?:(\d+):\d+: ([a-zA-Z\-/]*): (.*)`)
	warningDirectiveRE = regexp.MustCompile(`^\s*/// (.*)`)
)

func TestChecker(t *testing.T) {
	// b := make([]byte, 1000)
	// var memLog = bytes.NewBuffer(b)
	// log.SetOutput(memLog)
	// file, err := os.Create("myfile")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// mw := io.MultiWriter(os.Stdout, file)
	// fmt.Fprintln(mw, "This line will be written to stdout and also to a file")

	for _, rule := range RuleList()[:] {
		t.Run(rule.Name(), func(t *testing.T) {
			if rule.Name() != "evalOrder" {
				return
			}
			pkgPath := testdataPkgPath + rule.Name()

			prog := newProg(t, pkgPath)
			files := prog.Imported[pkgPath].Files
			ctx := NewContext(prog.Fset, sizes)

			for _, f := range files {
				filename := getFilename(prog, f)
				testFilename := filepath.Join("testdata", rule.Name(), filename)
				testWarns := newTestFileWarning(t, testFilename)

				warns := NewChecker(rule, ctx).Check(f)
				t.Logf("scanned %v, expected %v", len(warns), len(testWarns))

				// if !warningRE.MatchString(l) {
				// 	// Something that doen't look like a warning.
				// 	// Probably debug output or checker runtime error.
				// 	unexpectedLines = append(unexpectedLines, l)
				// 	continue
				// }

				var unexpectedLines []string

				for _, warn := range warns {

					var lineString, ruleName, text string
					unpackSubmatches(warn.Text, warningRE, &lineString, &ruleName, &text)
					line, err := strconv.Atoi(lineString)
					if err != nil {
						t.Errorf("%s: invalid line number in %s", testFilename, lineString)
					}
					if ruleName != rule.name {
						t.Errorf("%s: unexpected checker name: %s", testFilename, ruleName)
						continue
					}
					if w := find(testWarns, line, text); w != nil {
						if w.matched {
							t.Errorf("%s:%d: multiple matches for %s", testFilename, line, w)
						}
						w.matched = true
					} else {
						unexpectedLines = append(unexpectedLines, warn.Text)
					}
				}

				for line := range testWarns {
					for _, w := range testWarns[line] {
						if w.matched {
							continue
						}
						t.Errorf("%s:%d: unmatched %s", testFilename, line, w)
					}
				}
				for _, l := range unexpectedLines {
					t.Errorf("unexpected line in output: %s", l)
				}

				// unexpectedLines := []string{}
				// for i, warn := range warns {
				// 	line := getWarnLine(ctx, warn.Node)
				// 	twarn, ok := testWarns[line]
				// 	if !ok {
				// 		unexpectedLines = append(unexpectedLines, warn.Text)
				// 		continue
				// 	}
				// 	if warn.Text != twarn {
				// 		t.Errorf("got `%s`, want `%s`", warn.Text, testWarns[i])
				// 	} else {
				// 		delete(testWarns, line)
				// 	}
				// }
			}
		})
	}
}

func find(warnings map[int][]*warning, line int, text string) *warning {
	for _, y := range warnings[line] {
		if text == y.text {
			return y
		}
	}
	return nil
}

func getWarnLine(ctx *Context, node ast.Node) int {
	loc := ctx.FileSet().Position(node.Pos()).String()
	num, _ := strconv.Atoi(strings.Split(loc, ":")[1])
	return num
}

func getFilename(prog *loader.Program, f *ast.File) string {
	// see https://github.com/golang/go/issues/24498
	fname := prog.Fset.Position(f.Pos()).String() // ex: /usr/go/src/pkg/main.go:1:1
	fname = strings.Split(fname, ":")[0]          // ex: /usr/go/src/pkg/main.go
	return filepath.Base(fname)                   // ex: main.go
}

func newProg(t *testing.T, pkgPath string) *loader.Program {
	conf := loader.Config{
		ParserMode: parser.ParseComments,
		TypeChecker: types.Config{
			Sizes: sizes,
		},
	}
	if _, err := conf.FromArgs([]string{pkgPath}, true); err != nil {
		t.Fatalf("resolve packages: %v", err)
	}
	prog, err := conf.Load()
	if err != nil {
		t.Fatal(err)
	}
	return prog
}

type warning struct {
	matched bool
	text    string
}

func newTestFileWarning(t *testing.T, filapath string) map[int][]*warning {
	testData, err := ioutil.ReadFile(filapath)
	if err != nil {
		t.Fatalf("can't find checker tests: %v", err)
	}
	lines := strings.Split(string(testData), "\n")

	warnings := make(map[int][]*warning)
	var pending []*warning
	for i, l := range lines {
		if !warningDirectiveRE.MatchString(l) {
			if len(pending) != 0 {
				line := i + 1
				warnings[line] = append([]*warning{}, pending...)
				pending = pending[:0]
			}
			continue
		}
		var m warning
		unpackSubmatches(l, warningDirectiveRE, &m.text)
		pending = append(pending, &m)
	}
	return warnings
}

func (w warning) String() string {
	return "/// " + w.text
}

func unpackSubmatches(s string, re *regexp.Regexp, dst ...*string) {
	submatches := re.FindStringSubmatch(s)
	// Skip [0] which is a "whole match".
	for i, submatch := range submatches[1:] {
		*dst[i] = submatch
	}
}
