package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func main() {
	srcRoot := flag.String("src-root", "",
		`path to directory that should be linted recursively`)
	enable := flag.String("enable", "all",
		`forwarded to linter "as is"`)
	linter := flag.String("linter", "kfulint",
		`linter command name used for it's execution`)
	exclude := flag.String("exclude", "testdata/|vendor/|builtin/",
		`regexp used to skip package names`)

	flag.Parse()

	if *srcRoot == "" {
		log.Fatal("empty -src-root")
	}
	*srcRoot = filepath.Clean(*srcRoot)

	excludeRE, err := regexp.Compile(*exclude)
	if err != nil {
		log.Fatalf("bad -exclude pattern: %v", err)
	}

	packages := make(map[string]bool)
	err = filepath.Walk(*srcRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		path = path[len(*srcRoot)+len("/"):]
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
	args = append(args, "-enable", *enable)
	for pkg := range packages {
		args = append(args, pkg)
	}
	cmd := exec.Command(*linter, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("lint error: %v", err)
	}
}
