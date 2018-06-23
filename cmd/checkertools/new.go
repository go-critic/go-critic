package checkertools

import (
	"bytes"
	"flag"
	"go/build"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// TODO: add more such tools, like checker-rename to easily rename
// checker type, update comments and rename testdata dir?

// NewMain implements gocritic sub-command entry point.
func NewMain() {
	flag.Usage = func() {
		log.Printf("usage: rulename")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	// TODO: add camel case check.
	ruleName := flag.Arg(0)

	// We expect gocritic sources to be under canonical import
	// path under GOPATH.
	root := filepath.Join(build.Default.GOPATH,
		"src", "github.com", "go-critic", "go-critic")

	testDir := createTestDir(ruleName, root)
	createPositiveTests(testDir)
	createNegativeTests(testDir)
	createCheckerFile(ruleName, root)
}

func writeFile(filename string, data []byte) error {
	log.Printf("%s: new file", filename)
	return ioutil.WriteFile(filename, data, 0666)
}

func makeDir(filename string) error {
	log.Printf("%s: new dir", filename)
	return os.Mkdir(filename, 0755)
}

func createTestDir(ruleName, root string) (testDir string) {
	testDir = filepath.Join(root, "cmd", "gocritic", "testdata", ruleName)
	if err := makeDir(testDir); err != nil {
		log.Fatalf("create test dir: %v", err)
	}
	return testDir
}

func createPositiveTests(testDir string) {
	const data = `package checker_test

// TODO: write tests that trigger newly added checker.

/// this is example of warning directive. It fails your tests right now
func stub() {}
`
	filename := filepath.Join(testDir, "positive_tests.go")
	if err := writeFile(filename, []byte(data)); err != nil {
		log.Fatalf("create positive tests: %v", err)
	}
}

func createNegativeTests(testDir string) {
	const data = `package checker_test

// TODO: write tests that do not trigger newly added checker.
`
	filename := filepath.Join(testDir, "negative_tests.go")
	if err := writeFile(filename, []byte(data)); err != nil {
		log.Fatalf("create negative tests: %v", err)
	}
}

func createCheckerFile(ruleName, root string) {
	tmpl := template.Must(template.New("").Parse(`
package lint

//! This is {{.TypeName}} one-line summary stub.
//
// Optional section for details.
//
// Before:
// badCode := "minimal example of code that triggers a warning"
//
// After:
// goodCode := "example of intended fix"

func init() {
	addChecker(&{{.TypeName}})
}

type {{.TypeName}} struct {
	checkerBase
}

// TODO: fix checker documentation comment.
// TODO: implement one of the interfaces from checker_base.go.
`))
	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]interface{}{
		"TypeName": ruleName + "Checker",
	})

	formattedSrc, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("gofmt error: %v", err)
	}

	filename := filepath.Join(root, "lint", ruleName+"_checker.go")
	if err := writeFile(filename, formattedSrc); err != nil {
		log.Fatalf("create checker file: %v", err)
	}
}
