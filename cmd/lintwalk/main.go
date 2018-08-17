package lintwalk

import (
	"flag"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/go-critic/go-critic/lint/flagparser"
)

func packagePath() []string {
	goPath := strings.Split(build.Default.GOPATH, string(os.PathListSeparator))
	return append([]string{runtime.GOROOT()}, goPath...)
}

func getPackagePrefix(dir string) string {
	for _, p := range packagePath() {
		if strings.HasPrefix(dir, p) {
			if res, err := filepath.Rel(filepath.Join(p, "src"), dir); err == nil {
				return res
			}
		}
	}
	return ""
}

func dirIsHidden(dir string) bool {
	return strings.HasPrefix(dir, ".") || strings.HasPrefix(dir, "_")
}

// Main implements gocritic sub-command entry point.
func Main() {
	flags := flagparser.NewFlagParser()

	exclude := flag.String("exclude", `testdata/|vendor/|builtin/`,
		`regexp used to skip package names`)
	checkHidden := flag.Bool("checkHidden", false,
		`whether to visit dirs those name start with "." or "_"`)

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatalf("expected exactly one project root argument")
	}
	srcRoot := flag.Arg(0)
	if srcRoot == "" {
		log.Fatal("empty -src-root")
	}
	srcRoot = filepath.Clean(srcRoot)

	srcRoot, err := filepath.Abs(srcRoot)

	if err != nil {
		log.Fatal(err)
	}

	excludeRE, err := regexp.Compile(*exclude)

	if err != nil {
		log.Fatalf("bad -exclude pattern: %v", err)
	}

	packages := map[string]bool{}

	err = filepath.Walk(srcRoot, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			if info == nil {
				return e
			}
			log.Printf("walk error: %v", e)
		}
		if info.IsDir() && dirIsHidden(filepath.Base(path)) {
			if *checkHidden {
				return nil // OK, visit anyway
			}
			return filepath.SkipDir
		}
		if !strings.HasSuffix(path, ".go") || excludeRE.MatchString(path) {
			return nil
		}

		importPath := getPackagePrefix(filepath.Dir(path))
		packages[importPath] = true
		return nil
	})

	if err != nil {
		log.Fatalf("walk src-root: %v", err)
	}

	args := append([]string{"check-package"}, flags.Args()...)

	for p := range packages {
		args = append(args, p)
	}

	/* #nosec */
	cmd := exec.Command("gocritic", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("lint error: %v", err)
	}
}
