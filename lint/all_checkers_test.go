package lint

import (
	"fmt"
	"go/ast"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
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

// testCfg is a config used to initialize checkers for end2end testing.
// The options should make checkers as aggressive as possible, making them
// match 100% of cases they potentially could.
var testCfg = map[string]map[string]interface{}{
	"captLocal": {"checkLocals": true},
}

func TestSanity(t *testing.T) {
	saneRules := ruleList[:0]

	for _, rule := range ruleList {
		t.Run(rule.Name(), func(t *testing.T) {
			pkgPath := testdataPkgPath + "/_sanity"

			pkg := newPackage(t, pkgPath)

			ctx := NewContext(pkg.Fset, sizes)
			ctx.SetPackageInfo(pkg.TypesInfo, pkg.Types)
			files := pkg.Syntax

			for _, f := range files {
				defer func() {
					r := recover()
					if r != nil {
						t.Errorf("unexpected panic: `%v`\n%s", r, debug.Stack())
					} else {
						saneRules = append(saneRules, rule)
					}
				}()

				_ = NewChecker(rule, ctx, testCfg[rule.Name()]).Check(f)
			}
		})
	}

	ruleList = saneRules
}

func TestCheckers(t *testing.T) {
	for _, rule := range ruleList {
		t.Run(rule.Name(), func(t *testing.T) {
			testRule := rule
			if testing.CoverMode() == "" {
				t.Parallel()
			}
			pkgPath := testdataPkgPath + testRule.Name()

			pkg := newPackage(t, pkgPath)
			ctx := NewContext(pkg.Fset, sizes)
			ctx.SetPackageInfo(pkg.TypesInfo, pkg.Types)

			checkFiles(t, testRule, ctx, pkg, pkgPath)
		})
	}
}

func checkFiles(t *testing.T, rule *Rule, ctx *Context, pkg *packages.Package, pkgPath string) {
	files := pkg.Syntax

	for _, f := range files {
		filename := getFilename(pkg, f)
		testFilename := filepath.Join("testdata", rule.Name(), filename)
		goldenWarns := newGoldenFile(t, testFilename)

		c := NewChecker(rule, ctx, testCfg[rule.Name()])
		stripDirectives(f)
		ctx.SetFileInfo(filename, f)

		for _, warn := range c.Check(f) {
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

func getFilename(pkg *packages.Package, f *ast.File) string {
	// see https://github.com/golang/go/issues/24498
	return filepath.Base(pkg.Fset.Position(f.Pos()).Filename)
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

func newPackage(t *testing.T, pkgPath string) *packages.Package {
	conf := &packages.Config{Mode: packages.LoadSyntax}
	pkgs, err := packages.Load(conf, pkgPath)

	if err != nil {
		t.Fatal(err)
	}
	if len(pkgs) != 1 {
		t.Fatalf("more than 1 packages loaded for %s", pkgPath)
	}

	return pkgs[0]
}
