package checkers

import (
	"bytes"
	"go/ast"
	"go/token"
	"io/ioutil"
	"log"

	"github.com/go-critic/go-critic/framework/linter"
	"github.com/quasilyte/go-ruleguard/ruleguard"
	"fmt"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "ruleguard"
	info.Tags = []string{"style", "experimental"}
	info.Params = linter.CheckerParams{
		"rules": {
			Value: "",
			Usage: "path to a gorules file",
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
	c := &ruleguardChecker{ctx: ctx}

	var rules []string
	if info.Params.IsStringSlice("rules") {
		rules = info.Params.StringSlice("rules")
	} else if info.Params.IsString("rules") {
		rules = []string{info.Params.String("rules")}
	} else {
		return c
	}

	for _, rulesFileName := range rules {
		if rulesFileName == "" {
			continue
		}
		rset, err := parseRuleGuardFile(rulesFileName)
		if err != nil {
			log.Printf("skipped file %s: unable to parse: %+v", rulesFileName, err)
			continue
		}
		c.rset = append(c.rset, goRuleSetWithFileName{
			RuleFileName: rulesFileName,
			GoRuleSet:    rset,
		})
	}
	return c
}

func parseRuleGuardFile(rulesFilename string) (*ruleguard.GoRuleSet, error) {
	// TODO(quasilyte): handle initialization errors better when we make
	// a transition to the go/analysis framework.
	//
	// For now, we log error messages and return a ruleguard checker
	// with an empty rules set.
	data, err := ioutil.ReadFile(rulesFilename)
	if err != nil {
		return nil, fmt.Errorf("ruleguard init error: %w", err)
	}

	fset := token.NewFileSet()
	rset, err := ruleguard.ParseRules(rulesFilename, fset, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("ruleguard init error: %w",err)
	}
	return rset, err
}

type goRuleSetWithFileName struct {
	RuleFileName string
	*ruleguard.GoRuleSet
}

type ruleguardChecker struct {
	ctx *linter.CheckerContext
	rset []goRuleSetWithFileName
}

func (c *ruleguardChecker) WalkFile(f *ast.File) {
	if len(c.rset) == 0 {
		return
	}

	for _, rset := range c.rset {
		ctx := &ruleguard.Context{
			Pkg:   c.ctx.Pkg,
			Types: c.ctx.TypesInfo,
			Sizes: c.ctx.SizesInfo,
			Fset:  c.ctx.FileSet,
			Report: func(n ast.Node, msg string, _ *ruleguard.Suggestion) {
				// TODO(quasilyte): investigate whether we should add a rule name as
				// a message prefix here.
				c.ctx.Warn(n, "%s: %v", rset.RuleFileName, msg)
			},
		}


		err := ruleguard.RunRules(ctx, f, rset.GoRuleSet)
		if err != nil {
			// Normally this should never happen, but since
			// we don't have a better mechanism to report errors,
			// emit a warning.
			c.ctx.Warn(f, "execution error: %s: %v", rset.RuleFileName, err)
		}
	}


}
