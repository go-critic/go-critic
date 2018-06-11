package lintwalk

import (
	"flag"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func packagePath() []string {
	res := []string{runtime.GOROOT(), build.Default.GOPATH}

	if p, ok := os.LookupEnv("GOPATH"); ok {
		res = append(res, strings.Split(p, ":")...)
	}
	return res
}

func getPackagePrefix(dir string) string {
	for _, p := range packagePath() {
		if filepath.HasPrefix(dir, p) {
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
	exclude := flag.String("exclude", "(.*test.*)|(.*vendor.*)|(.*builtin.*)",
		`regexp used to skip package names`)

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

	testRE := regexp.MustCompile(".*_test")

	if err != nil {
		log.Fatal(err)
	}

	fs := token.NewFileSet()

	packages := map[string]bool{}

	err = filepath.Walk(srcRoot, func(path string, info os.FileInfo, err error) error {
		p, err := filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
		}
		pkgs, err := parser.ParseDir(fs, p, nil, parser.PackageClauseOnly)

		if !info.IsDir() || excludeRE.MatchString(p) {
			return nil
		}
		p = filepath.Dir(p)
		p = getPackagePrefix(p)

		for _, pk := range pkgs {
			if testRE.MatchString(pk.Name) {
				continue
			}
			if pk.Name == "main" {
				packages[p] = true
			} else {
				packages[filepath.Join(p, pk.Name)] = true
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	var args []string
	args = append(args, "check-package", "-enable", *enable)

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
