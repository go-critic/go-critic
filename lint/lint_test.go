package lint

import (
	"go/ast"
	"go/parser"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"golang.org/x/tools/go/loader"
)

var (
	testdataPkgPath    = "github.com/go-critic/go-critic/lint/testdata/"
	sizes              = types.SizesFor("gc", runtime.GOARCH)
	warningDirectiveRE = regexp.MustCompile(`^\s*/// (.*)`)
)

var ruleList []*Rule

func TestMain(m *testing.M) {
	// Can't do RuleList call in ordinary init due to initialization
	// order dependency. There are 2 solutions:
	//	1. make this package "external" test like "lint_test"
	//	2. use TestMain to run tests after we assign ruleList, after inits.
	// We're going with (2) here.
	ruleList = RuleList()
	os.Exit(m.Run())
}

func TestSanity(t *testing.T) {
	saneRules := ruleList[:0]

	for _, rule := range ruleList {
		t.Run(rule.Name(), func(t *testing.T) {
			pkgPath := testdataPkgPath + "/_sanity"

			prog := newProg(t, pkgPath)
			pkgInfo := prog.Imported[pkgPath]

			ctx := NewContext(prog.Fset, sizes)
			ctx.SetPackageInfo(&pkgInfo.Info, pkgInfo.Pkg)

			files := prog.Imported[pkgPath].Files

			for _, f := range files {
				defer func() {
					r := recover()
					if r != nil {
						t.Errorf("unexpected panic: `%v`", r)
					} else {
						saneRules = append(saneRules, rule)
					}
				}()

				_ = NewChecker(rule, ctx, nil).Check(f)
			}
		})
	}

	ruleList = saneRules
}

func TestCheckers(t *testing.T) {
	for _, rule := range ruleList {
		t.Run(rule.Name(), func(t *testing.T) {
			pkgPath := testdataPkgPath + rule.Name()

			prog := newProg(t, pkgPath)
			pkgInfo := prog.Imported[pkgPath]

			ctx := NewContext(prog.Fset, sizes)
			ctx.SetPackageInfo(&pkgInfo.Info, pkgInfo.Pkg)

			checkFiles(t, rule, ctx, prog, pkgPath)
		})
	}
}

func checkFiles(t *testing.T, rule *Rule, ctx *Context, prog *loader.Program, pkgPath string) {
	files := prog.Imported[pkgPath].Files

	for _, f := range files {
		filename := getFilename(prog, f)
		testFilepath := filepath.Join("testdata", rule.Name(), filename)
		goldenWarns := newGoldenFile(t, testFilepath)

		var unexpectedWarns []string

		warns := NewChecker(rule, ctx, nil).Check(f)

		for _, warn := range warns {
			line := ctx.FileSet().Position(warn.Node.Pos()).Line

			if w := goldenWarns.find(line, warn.Text); w != nil {
				if w.matched {
					t.Errorf("%s:%d: multiple matches for %s", testFilepath, line, w)
				}
				w.matched = true
			} else {
				unexpectedWarns = append(unexpectedWarns, warn.Text)
			}
		}

		goldenWarns.checkUnmatched(t, testFilepath)

		for _, l := range unexpectedWarns {
			t.Errorf("unexpected warn: `%s`", l)
		}
	}
}

func getFilename(prog *loader.Program, f *ast.File) string {
	// see https://github.com/golang/go/issues/24498
	return filepath.Base(prog.Fset.Position(f.Pos()).Filename)
}

type goldenFile struct {
	warnings map[int][]*warning
}

type warning struct {
	matched bool
	text    string
}

func (w warning) String() string {
	return w.text
}

func newGoldenFile(t *testing.T, filepath string) *goldenFile {
	testData, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatalf("can't find checker tests: %v", err)
	}
	lines := strings.Split(string(testData), "\n")

	warnings := make(map[int][]*warning)
	var pending []*warning

	for i, l := range lines {
		if warningDirectiveRE.MatchString(l) {
			var m warning
			unpackSubmatches(l, warningDirectiveRE, &m.text)
			pending = append(pending, &m)
		} else {
			if len(pending) != 0 {
				line := i + 1
				warnings[line] = append([]*warning{}, pending...)
				pending = pending[:0]
			}
		}
	}
	return &goldenFile{warnings: warnings}
}

func (f *goldenFile) find(line int, text string) *warning {
	for _, y := range f.warnings[line] {
		if text == y.text {
			return y
		}
	}
	return nil
}

func (f *goldenFile) checkUnmatched(t *testing.T, testFilepath string) {
	for line := range f.warnings {
		for _, w := range f.warnings[line] {
			if w.matched {
				continue
			}
			t.Errorf("%s:%d: unmatched `%s`", testFilepath, line, w)
		}
	}
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
	pkgInfo := prog.Imported[pkgPath]
	if pkgInfo == nil || !pkgInfo.TransitivelyErrorFree {
		t.Fatalf("%s package is not properly loaded", pkgPath)
	}
	return prog
}

func unpackSubmatches(s string, re *regexp.Regexp, dst ...*string) {
	submatches := re.FindStringSubmatch(s)
	// Skip [0] which is a "whole match".
	if len(submatches) > 0 {
		for i, submatch := range submatches[1:] {
			*dst[i] = submatch
		}
	}
}
