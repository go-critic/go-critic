package checkers

import (
	"testing"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/linttest"
)

func TestCheckers(t *testing.T) {
	allParams := map[string]map[string]interface{}{
		"captLocal": {"paramsOnly": false},
	}

	for _, info := range lintpack.GetCheckersInfo() {
		params := allParams[info.Name]
		for key, p := range info.Params {
			v, ok := params[key]
			if ok {
				p.Value = v
			}
		}
	}

	linttest.TestCheckers(t)
}

func TestIntegration(t *testing.T) { linttest.TestIntegration(t) }

func TestTags(t *testing.T) {
	// Verify that we're only using strict set of tags.
	// This helps to avoid typos in tag names.
	//
	// Also check that exactly 1 category tag is used.

	for _, info := range lintpack.GetCheckersInfo() {
		categories := 0
		for _, tag := range info.Tags {
			switch tag {
			case "diagnostic", "style", "performance":
				// Category tags.
				// Can only have one of them.
				categories++
			case "experimental", "opinionated":
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
	for _, info := range lintpack.GetCheckersInfo() {
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
	// Verify that new checker is not added without "experimental"
	// tag by accident. When stable checker is about to be added,
	// slice above should be modified to include new checker name.

	// It is a good practice to keep this list sorted.
	stableList := []string{
		"appendAssign",
		"appendCombine",
		"assignOp",
		"builtinShadow",
		"captLocal",
		"caseOrder",
		"defaultCaseOrder",
		"dupArg",
		"dupBranchBody",
		"dupCase",
		"elseif",
		"flagDeref",
		"ifElseChain",
		"importShadow",
		"indexAlloc",
		"paramTypeCombine",
		"rangeExprCopy",
		"rangeValCopy",
		"regexpMust",
		"singleCaseSwitch",
		"sloppyLen",
		"switchTrue",
		"typeSwitchVar",
		"typeUnparen",
		"underef",
		"unlambda",
		"unslice",
		"dupSubExpr",
		"hugeParam",
	}

	m := make(map[string]bool)
	for _, name := range stableList {
		m[name] = true
	}

	for _, info := range lintpack.GetCheckersInfo() {
		if info.HasTag("experimental") {
			continue
		}
		if !m[info.Name] {
			t.Errorf("%q checker misses `experimental` tag", info.Name)
		}
	}
}
