package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"runtime"
	"testing"

	"github.com/go-critic/go-critic/linter"
)

func TestShortenLocation(t *testing.T) {
	testGopath := "/home/queen/go/"
	testGoroot := "/usr/lib/go/"
	tests := []struct {
		wd    string
		input string
		out   string
	}{
		{"", "/home/queen/go/file.go", "$GOPATH/file.go"},
		{"", "/home/queen/go-file.go", "/home/queen/go-file.go"},

		{"", "/usr/lib/go/file.go", "$GOROOT/file.go"},
		{"", "/usr/lib/go-file.go", "/usr/lib/go-file.go"},

		{"/home/queen/go/src/", "/home/queen/go/src/file.go", "./file.go"},
		{"/home/queen/", "/home/queen-src-file.go", "/home/queen-src-file.go"},
		{"/home/", "/home/queen/go/src/file.go", "$GOPATH/src/file.go"},

		{`C:\home\queen\go\src\`, `C:\home\queen\go\src\file.go`, "./file.go"},
	}

	l := &program{
		gopath: testGopath,
		goroot: testGoroot,
	}
	for _, test := range tests {
		l.workDir = test.wd
		have := l.shortenLocation(test.input)
		want := test.out
		if have != want {
			t.Errorf("shorten(%q):\nhave: %q\nwant: %q",
				test.input, have, want)
		}
	}
}

// panickingWalker is a FileWalker that panics when WalkFile is called.
type panickingWalker struct {
	err any
}

func (w *panickingWalker) WalkFile(_ *ast.File) {
	panic(w.err)
}

func TestCheckFile_PanicRecovery(t *testing.T) {
	// Register a checker that panics with an error.
	collection := &linter.CheckerCollection{}
	errToPanic := errors.New("test checker panic")

	collection.AddChecker(&linter.CheckerInfo{
		Name: "testPanicChecker",
		Tags: []string{"diagnostic"},
	}, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return &panickingWalker{err: errToPanic}, nil
	})

	// Create a linter context and checker.
	fset := token.NewFileSet()
	sizes := types.SizesFor("gc", runtime.GOARCH)
	ctx := linter.NewContext(fset, sizes)
	info := linter.GetCheckersInfo()

	var checkerInfo *linter.CheckerInfo
	for _, i := range info {
		if i.Name == "testPanicChecker" {
			checkerInfo = i
			break
		}
	}
	if checkerInfo == nil {
		t.Fatal("testPanicChecker not registered")
	}

	checker, err := linter.NewChecker(ctx, checkerInfo)
	if err != nil {
		t.Fatalf("NewChecker: %v", err)
	}

	// Parse a minimal Go file.
	src := `package main
func main() {}
`
	f, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}

	// Create a program and verify that checkFile re-throws the panic.
	p := &program{
		ctx:        ctx,
		checkers:   []*linter.Checker{checker},
		concurrency: runtime.GOMAXPROCS(0),
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic from checkFile, but none occurred")
		}
		if r != errToPanic {
			t.Errorf("recovered panic = %v (%T), want %v", r, r, errToPanic)
		}
	}()

	p.checkFile(f)

	// Should not reach here — panic should propagate.
	t.Fatal("checkFile returned without re-throwing panic")
}
