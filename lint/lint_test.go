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
	warningDirectiveRE = regexp.MustCompile(`^\s*/// (.*)`)
)

func TestChecker(t *testing.T) {
	for _, rule := range RuleList() {
		t.Run(rule.Name(), func(t *testing.T) {
			pkgPath := testdataPkgPath + rule.Name()

			prog := newProg(t, pkgPath)
			pkgInfo := prog.Imported[pkgPath]
			files := prog.Imported[pkgPath].Files

			ctx := NewContext(prog.Fset, sizes)
			ctx.SetPackageInfo(&pkgInfo.Info, pkgInfo.Pkg)

			for _, f := range files {
				filename := getFilename(prog, f)
				testFilepath := filepath.Join("testdata", rule.Name(), filename)
				goldenWarns := newGoldenFile(t, testFilepath)

				var unexpectedWarns []string

				warns := NewChecker(rule, ctx).Check(f)

				for _, warn := range warns {
					line := getWarnLine(ctx, warn.Node)

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
		})
	}
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
