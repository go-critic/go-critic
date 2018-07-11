package lint

import (
	"fmt"
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
	commentRE          = regexp.MustCompile(`^\s*//`)
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
			if testing.CoverMode() == "" {
				t.Parallel()
			}
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
		testFilename := filepath.Join("testdata", rule.Name(), filename)
		goldenWarns := newGoldenFile(t, testFilename)

		stripDirectives(f)
		warns := NewChecker(rule, ctx, nil).Check(f)

		for _, warn := range warns {
			line := ctx.FileSet().Position(warn.Node.Pos()).Line

			if w := goldenWarns.find(line, warn.Text); w != nil {
				if w.matched {
					t.Errorf("%s:%d: multiple matches for %s",
						testFilename, line, w)
				}
				w.matched = true
			} else {
				t.Errorf("%s:%d: unexpected warn: %s",
					testFilename, line, warn.Text)
			}
		}

		goldenWarns.checkUnmatched(t, testFilename)
	}
}

// stripDirectives replaces "///" comments with empty single-line
// comments, so the checkers that inspect comments see ordinary
// comment groups (with extra newlines, but that's not important).
func stripDirectives(f *ast.File) {
	for _, cg := range f.Comments {
		for _, c := range cg.List {
			if strings.HasPrefix(c.Text, "/// ") {
				c.Text = "//"
			}
		}
	}
}

func TestIncorrectRule(t *testing.T) {
	func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected panic on nil rule")
			}
			if r != "nil rule given" {
				t.Fatalf("expected `nil rule given`, got %v", r)
			}
		}()
		NewChecker(nil, nil, nil)
	}(t)

	func(t *testing.T) {
		name := "i-don-not-exist"

		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected panic on incorrect name")
			}

			want := fmt.Sprintf("rule %q is undefined", name)
			if r != want {
				t.Fatalf("expected `%v`, got %v", want, r)
			}
		}()

		r := &Rule{name: name}
		NewChecker(r, nil, nil)
	}(t)
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

func newGoldenFile(t *testing.T, filename string) *goldenFile {
	testData, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("can't find checker tests: %v", err)
	}
	lines := strings.Split(string(testData), "\n")

	warnings := make(map[int][]*warning)
	var pending []*warning

	for i, l := range lines {
		if m := warningDirectiveRE.FindStringSubmatch(l); m != nil {
			pending = append(pending, &warning{text: m[1]})
		} else if len(pending) != 0 {
			line := i + 1
			if commentRE.MatchString(l) {
				// Hack to make it possible to attach directives
				// to a proper single-line comment position.
				line -= len(pending)
			}
			warnings[line] = append([]*warning{}, pending...)
			pending = pending[:0]
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

func (f *goldenFile) checkUnmatched(t *testing.T, testFilename string) {
	for line := range f.warnings {
		for _, w := range f.warnings[line] {
			if w.matched {
				continue
			}
			t.Errorf("%s:%d: unmatched `%s`", testFilename, line, w)
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
