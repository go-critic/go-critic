package linttest

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Run executes integration tests.
func (cfg *IntegrationTest) Run(t *testing.T) {
	absDir, err := filepath.Abs(cfg.Dir)
	if err != nil {
		t.Fatalf("can't get dir abs path: %v", err)
	}

	gocritic, err := cfg.buildLinter()
	if err != nil {
		t.Fatalf("build linter: %v", err)
	}

	files, err := os.ReadDir(absDir)
	if err != nil {
		t.Fatalf("list test files: %v", err)
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		t.Run(f.Name(), func(t *testing.T) {
			wd := filepath.Join(absDir, f.Name())
			if err := os.Chdir(wd); err != nil {
				t.Fatalf("enter test dir: %v", err)
			}
			cfg.runTest(t, gocritic, wd)
		})
	}
}

func (cfg *IntegrationTest) runTest(t *testing.T, gocritic, gopath string) {
	data, err := os.ReadFile("linttest.params")
	if err != nil {
		t.Fatalf("reading linter run params: %v", err)
	}

	// If several tests re-use a single golden file,
	// don't read it repeatedly, just re-use its contents.
	goldenDataCache := make(map[string]string)

	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		// The format is:
		//	runParams ... "|" goldenFile
		parts := strings.Split(line, "|")
		runParams := strings.Split(strings.TrimSpace(parts[0]), " ")
		goldenFile := strings.TrimSpace(parts[1])

		// Read from a golden file or contents cache.
		var want string
		if data, ok := goldenDataCache[goldenFile]; ok {
			want = data
		} else {
			data, err := os.ReadFile(goldenFile)
			if err != nil {
				t.Errorf("read golden file: %v", err)
			}
			want = string(bytes.TrimSpace(data))
			goldenDataCache[goldenFile] = want
		}

		// Get the actual execution output.
		cmd := exec.Command(gocritic, runParams...)
		cmd.Env = append([]string{}, os.Environ()...) // Copy parent env
		cmd.Env = append(cmd.Env,
			// Override GOPATH.
			"GOPATH="+gopath,
			"GO111MODULE=auto")

		out, err := cmd.CombinedOutput()
		out = bytes.TrimSpace(out)
		var have string
		if err != nil {
			// Error is prepended to the beginning.
			have = err.Error() + "\n" + string(out)
		} else {
			have = string(out)
		}

		// To get line-by-line diff, split is required.
		wantLines := strings.Split(want, "\n")
		haveLines := strings.Split(have, "\n")
		if diff := cmp.Diff(wantLines, haveLines); diff != "" {
			t.Errorf("linttest.params:%d: output mismatch:\n%s", i+1, diff)
			t.Logf("linter output was: %s\n", have)
		}
	}
}

func (cfg *IntegrationTest) buildLinter() (string, error) {
	tmpDir := os.TempDir()
	filename := filepath.Join(tmpDir, "_gocritic_inttest_")

	args := []string{"build", "-race", "-o", filename, cfg.Main}
	out, err := exec.Command("go", args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%v: %s", err, out)
	}

	return filename, nil
}
