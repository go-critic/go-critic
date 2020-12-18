package checkers

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-critic/go-critic/framework/linter"
	"github.com/quasilyte/go-ruleguard/ruleguard"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "ruleguard"
	info.Tags = []string{"style", "experimental"}
	info.Params = linter.CheckerParams{
		"rules": {
			Value: "",
			Usage: "comma-separated list of gorule file paths. Patterns such as 'rules-*.go' may be specified",
		},
		"debug": {
			Value: "",
			Usage: "enable debug for the specified named rules group",
		},
		"failOnError": {
			Value: false,
			Usage: "If true, panic when the gorule files contain a syntax error. If false, log and skip rules that contain an error",
		},
	}
	info.Summary = "Runs user-defined rules using ruleguard linter"
	info.Details = "Reads a rules file and turns them into go-critic checkers."
	info.Before = `N/A`
	info.After = `N/A`
	info.Note = "See https://github.com/quasilyte/go-ruleguard."

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) linter.FileWalker {
		return newRuleguardChecker(&info, ctx)
	})
}

func newRuleguardChecker(info *linter.CheckerInfo, ctx *linter.CheckerContext) *ruleguardChecker {
	c := &ruleguardChecker{
		ctx:        ctx,
		debugGroup: info.Params.String("debug"),
	}
	rulesFlag := info.Params.String("rules")
	if rulesFlag == "" {
		return c
	}
	failOnErrorFlag := info.Params.Bool("failOnError")

	// TODO(quasilyte): handle initialization errors better when we make
	// a transition to the go/analysis framework.
	//
	// For now, we log error messages and return a ruleguard checker
	// with an empty rules set.

	fset := token.NewFileSet()
	filePatterns := strings.Split(rulesFlag, ",")
	var ruleSets []*ruleguard.GoRuleSet
	for _, filePattern := range filePatterns {
		filenames, err := filepath.Glob(strings.TrimSpace(filePattern))
		if err != nil {
			// The only possible returned error is ErrBadPattern, when pattern is malformed.
			log.Printf("ruleguard init error: %+v", err)
			continue
		}
		for _, filename := range filenames {
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				if failOnErrorFlag {
					log.Panicf("ruleguard init error: %+v", err)
				}
				log.Printf("ruleguard init error: %+v", err)
				continue
			}
			rset, err := ruleguard.ParseRules(filename, fset, bytes.NewReader(data))
			if err != nil {
				if failOnErrorFlag {
					log.Panicf("ruleguard init error: %+v", err)
				}
				log.Printf("ruleguard init error: %+v", err)
				return c
			}
			ruleSets = append(ruleSets, rset)
		}
	}

	c.rset = ruleguard.MergeRuleSets(ruleSets)
	return c
}

type ruleguardChecker struct {
	ctx *linter.CheckerContext

	debugGroup string
	rset       *ruleguard.GoRuleSet
}

func (c *ruleguardChecker) WalkFile(f *ast.File) {
	if c.rset == nil {
		return
	}

	ctx := &ruleguard.Context{
		Debug: c.debugGroup,
		DebugPrint: func(s string) {
			fmt.Fprintln(os.Stderr, s)
		},
		Pkg:   c.ctx.Pkg,
		Types: c.ctx.TypesInfo,
		Sizes: c.ctx.SizesInfo,
		Fset:  c.ctx.FileSet,
		Report: func(_ ruleguard.GoRuleInfo, n ast.Node, msg string, _ *ruleguard.Suggestion) {
			// TODO(quasilyte): investigate whether we should add a rule name as
			// a message prefix here.
			c.ctx.Warn(n, msg)
		},
	}

	err := ruleguard.RunRules(ctx, f, c.rset)
	if err != nil {
		// Normally this should never happen, but since
		// we don't have a better mechanism to report errors,
		// emit a warning.
		c.ctx.Warn(f, "execution error: %v", err)
	}
}
