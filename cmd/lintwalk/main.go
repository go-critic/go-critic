package lintwalk

import (
	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func packagePath() []string {
	return []string{runtime.GOROOT(), build.Default.GOPATH}
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

// Main implements gocritic sub-command entry point.
func Main() {
	enable := flag.String("enable", "all",
		`forwarded to linter "as is"`)
	exclude := flag.String("exclude", "testdata/|vendor/|builtin/",
		`regexp used to skip package names`)
	checkGenerated := flag.Bool("checkGenerated", false, `forwarded to linter "as is"`)
	shorterErrLocation := flag.Bool("shorterErrLocation", true, `forwarded to linter "as is"`)

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
		if info.IsDir() || !strings.HasSuffix(path, ".go") || excludeRE.MatchString(path) {
			return nil
		}

		path = filepath.Dir(path)

		path = getPackagePrefix(path)

		packages[path] = true
		return nil
	})

	if err != nil {
		log.Fatalf("walk src-root: %v", err)
	}

	args := []string{
		"check-package",
		"-enable", *enable,
		"-checkGenerated=" + fmt.Sprint(*checkGenerated),
		"-shorterErrLocation=" + fmt.Sprint(*shorterErrLocation),
	}
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
