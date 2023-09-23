package checkers

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-critic/go-critic/checkers/internal/linttest"
	"github.com/go-critic/go-critic/linter"

	"github.com/google/go-cmp/cmp"
)

func init() {
	if err := InitEmbeddedRules(); err != nil {
		panic(err) // Should never happen
	}
}

func TestCheckers(t *testing.T) {
	allParams := map[string]map[string]interface{}{
		"captLocal":        {"paramsOnly": false},
		"commentedOutCode": {"minLength": 9},
	}

	for _, info := range linter.GetCheckersInfo() {
		params := allParams[info.Name]
		for key, p := range info.Params {
			v, ok := params[key]
			if ok {
				p.Value = v
			}
		}
	}

	cfg := linttest.CheckersTest{
		IgnoreErrors: []string{
			"caseOrder",
		},
	}

	cfg.Run(t)
}

func TestIntegration(t *testing.T) {
	cfg := linttest.IntegrationTest{
		Main: "github.com/go-critic/go-critic/cmd/gocritic",
		Dir:  "./testdata/_integration",
	}
	cfg.Run(t)
}

func TestTags(t *testing.T) {
	// Verify that we're only using strict set of tags.
	// This helps to avoid typos in tag names.
	//
	// Also check that exactly 1 category tag is used.

	for _, info := range linter.GetCheckersInfo() {
		categories := 0
		for _, tag := range info.Tags {
			switch tag {
			case linter.DiagnosticTag, linter.StyleTag, linter.PerformanceTag:
				// Category tags. Can only have one of them.
				categories++
			case linter.ExperimentalTag, linter.OpinionatedTag:
				// Optional tags.
			default:
				t.Errorf("%q checker uses unknown tag %q", info.Name, tag)
			}
		}
		if categories != 1 {
			t.Errorf("%q expected to have 1 category, found %d",
				info.Name, categories)
		}
	}
}

func TestDocs(t *testing.T) {
	for _, info := range linter.GetCheckersInfo() {
		if info.Summary == "" {
			t.Errorf("%q checker lacks summary", info.Name)
		}
		for key, p := range info.Params {
			if p.Usage == "" {
				t.Errorf("%q checker %q param lacks usage docs",
					info.Name, key)
			}
		}
	}
}

func TestStableList(t *testing.T) {
	// Verify that new checker is not added without linter.ExperimentalTag
	// tag by accident. When stable checker is about to be added,
	// slice above should be modified to include new checker name.

	// It is a good practice to keep this list sorted.
	stableList := []string{
		"appendAssign",
		"appendCombine",
		"argOrder",
		"assignOp",
		"badCall",
		"badCond",
		"builtinShadow",
		"captLocal",
		"caseOrder",
		"codegenComment",
		"commentFormatting",
		"defaultCaseOrder",
		"deprecatedComment",
		"dupArg",
		"dupBranchBody",
		"dupCase",
		"dupSubExpr",
		"elseif",
		"exitAfterDefer",
		"flagDeref",
		"flagName",
		"hugeParam",
		"ifElseChain",
		"importShadow",
		"indexAlloc",
		"mapKey",
		"newDeref",
		"offBy1",
		"paramTypeCombine",
		"rangeExprCopy",
		"rangeValCopy",
		"regexpMust",
		"singleCaseSwitch",
		"sloppyLen",
		"sloppyTypeAssert",
		"stringXbytes",
		"switchTrue",
		"typeSwitchVar",
		"typeUnparen",
		"underef",
		"unlambda",
		"unslice",
		"valSwap",
		"wrapperFunc",
	}

	m := make(map[string]bool)
	for _, name := range stableList {
		m[name] = true
	}

	for _, info := range linter.GetCheckersInfo() {
		if info.HasTag(linter.ExperimentalTag) {
			continue
		}
		if !m[info.Name] {
			t.Errorf("%q checker misses `experimental` tag", info.Name)
		}
	}
}

func TestExternal(t *testing.T) {
	t.Skip("temporary disabled during bump to Go 1.20")

	// Don't run these tests normally, unless asked to.
	// Note that CI tests do enable GOCRITIC_EXTERNAL_TESTS.
	if os.Getenv("GOCRITIC_EXTERNAL_TESTS") == "" {
		t.Skip("GOCRITIC_EXTERNAL_TESTS is unset")
	}

	// Build the linter.
	tmpDir := os.TempDir()
	gocriticBin := filepath.Join(tmpDir, "gocritic_external_test.exe")
	args := []string{"build", "-race", "-o", gocriticBin, "github.com/go-critic/go-critic/cmd/gocritic"}
	out, err := exec.Command("go", args...).CombinedOutput()
	if err != nil {
		t.Fatalf("%v: %s", err, out)
	}
	defer os.Remove(gocriticBin)

	externalTests := filepath.Join(tmpDir, "extern-testdata")
	out, err = exec.Command("git", "clone", "https://github.com/go-critic/extern-testdata.git", externalTests).CombinedOutput()
	if err != nil {
		t.Fatalf("%v: %s", err, out)
	}
	defer os.RemoveAll(externalTests)

	projects, err := os.ReadDir(filepath.Join(externalTests, "projects"))
	if err != nil {
		t.Fatal(err)
	}

	if len(projects) == 0 {
		t.Fatal("found 0 test projects")
	}

	for _, proj := range projects {
		projectPath := filepath.Join(externalTests, "projects", proj.Name())
		gocriticCmd := exec.Command(gocriticBin, "check", "--enableAll", "./...")
		gocriticCmd.Dir = projectPath
		out, err := gocriticCmd.CombinedOutput()
		have := err.Error() + "\n" + strings.TrimSpace(string(out))
		wantBytes, err := os.ReadFile(filepath.Join(projectPath, "output.golden"))
		want := strings.TrimSpace(string(wantBytes))
		if err != nil {
			t.Errorf("read %s golden file: %v", proj.Name(), err)
			continue
		}
		if diff := cmp.Diff(want, have); diff != "" {
			t.Errorf("%s output mismatches (+have -want):\n%s", proj.Name(), diff)
			continue
		}
	}
}
