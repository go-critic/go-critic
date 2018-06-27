package main_test

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

const (
	// binary is test linter executable name.
	// Using "exe" suffix to make it work on Windows as well.
	binary = "testlint.exe"
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
