package checkers

import (
	"bytes"
	"go/ast"
	"go/token"
	"io/ioutil"
	"log"

	"github.com/go-lintpack/lintpack"
	"github.com/quasilyte/go-ruleguard/ruleguard"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "ruleguard"
	info.Tags = []string{"style", "experimental"}
	info.Params = lintpack.CheckerParams{
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

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		return newRuleguardChecker(&info, ctx)
	})
}

func newRuleguardChecker(info *lintpack.CheckerInfo, ctx *lintpack.CheckerContext) *ruleguardChecker {
	c := &ruleguardChecker{ctx: ctx}
	rulesFilename := info.Params.String("rules")
	if rulesFilename == "" {
		return c
	}

	// TODO(quasilyte): handle initialization errors better when we make
	// a transition to the go/analysis framework.
	//
	// For now, we log error messages and return a ruleguard checker
	// with an empty rules set.

	data, err := ioutil.ReadFile(rulesFilename)
	if err != nil {
		log.Printf("ruleguard init error: %+v", err)
		return c
	}

	fset := token.NewFileSet()
	rset, err := ruleguard.ParseRules(rulesFilename, fset, bytes.NewReader(data))
	if err != nil {
		log.Printf("ruleguard init error: %+v", err)
		return c
	}

	c.rset = rset
	return c
}

type ruleguardChecker struct {
	ctx *lintpack.CheckerContext

	rset *ruleguard.GoRuleSet
}

func (c *ruleguardChecker) WalkFile(f *ast.File) {
	if c.rset == nil {
		return
	}

	ctx := &ruleguard.Context{
		Pkg:   c.ctx.Pkg,
		Types: c.ctx.TypesInfo,
		Sizes: c.ctx.SizesInfo,
		Fset:  c.ctx.FileSet,
		Report: func(n ast.Node, msg string, _ *ruleguard.Suggestion) {
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
