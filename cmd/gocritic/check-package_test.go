package main_test

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"testing"

	"github.com/go-critic/go-critic/lint"
)

const (
	// binary is test linter executable name.
	// Using "exe" suffix to make it work on Windows as well.
	binary = "testlint.exe"

	// linterCmdPath holds full path to linter main pkg path.
	linterCmdPath = "github.com/go-critic/go-critic/cmd/gocritic/"
)

func TestMain(m *testing.M) {
	// Before all tests are executed, we need to build linter first.
	// After all tests are completed, linter executable is removed.

	buildLinter()
	// Run all tests in a normal way.
	exitStatus := m.Run()
	// Nothing bad will happen if someone already removed the binary.
	_ = os.Remove(binary)

	os.Exit(exitStatus)
}

type ruleTest struct {
	name string
}

var ruleList []ruleTest

func init() {
	for _, rule := range lint.RuleList() {
		ruleList = append(ruleList, ruleTest{
			name: rule.Name(),
		})
	}
}

func runChecker(name, pkgPath string) (output []byte, err error) {
	cmd := exec.Command("./"+binary,
		"check-package",
		"-enable", name,
		"-failcode", "0",
		pkgPath)
	return cmd.CombinedOutput()
}

func TestSanity(t *testing.T) {
	saneRules := ruleList[:0]
	for _, rule := range ruleList {
		pkgPath := linterCmdPath + "testdata/_sanity"
		output, err := runChecker(rule.name, pkgPath)
		if err != nil {
			t.Errorf("%s failed sanity checks: %v:\n%s", rule.name, err, output)
			continue
		}
		saneRules = append(saneRules, rule)
	}
	ruleList = saneRules
}

func buildLinter() {
	// TODO(quasilyte): check that this actually works on windows.
	goBin := "go"
	if runtime.GOOS == "windows" {
		goBin += ".exe"
	}
	output, err := exec.Command(goBin, "build", "-o", binary).CombinedOutput()
	if err != nil {
		log.Fatalf("build linter: %v: %s", err, output)
	}
}
