package lintwalk

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

// Main implements gocritic sub-command entry point.
func Main() {
	enable := flag.String("enable", "all",
		`forwarded to linter "as is"`)
	exclude := flag.String("exclude", "testdata/|vendor/|builtin/",
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

	excludeRE, err := regexp.Compile(*exclude)
	if err != nil {
		log.Fatalf("bad -exclude pattern: %v", err)
	}

	packages := make(map[string]bool)
	err = filepath.Walk(srcRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("walk error: %v", err)
			return nil
		}
		if info.IsDir() {
			return nil
		}
		path = path[len(srcRoot)+len("/"):]
		if filepath.Ext(path) != ".go" {
			return nil
		}
		if excludeRE.MatchString(path) {
			return nil
		}
		packages[filepath.Dir(path)] = true
		return nil
	})
	if err != nil {
		log.Fatalf("walk src-root: %v", err)
	}

	var args []string
	args = append(args, "check-package", "-enable", *enable)
	for pkg := range packages {
		args = append(args, pkg)
	}
	/* #nosec */
	cmd := exec.Command("gocritic", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("lint error: %v", err)
	}
}
