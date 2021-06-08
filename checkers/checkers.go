// Package checkers is a gocritic linter main checkers collection.
package checkers

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"os"

	"github.com/go-critic/go-critic/checkers/rulesdata"
	"github.com/go-critic/go-critic/framework/linter"
	"github.com/quasilyte/go-ruleguard/ruleguard"
)

var collection = &linter.CheckerCollection{
	URL: "https://github.com/go-critic/go-critic/checkers",
}

var debug = func() func() bool {
	v := os.Getenv("DEBUG") != ""
	return func() bool {
		return v
	}
}()

// go-bindata is github.com/shuLhan/go-bindata;
// TODO(quasilyte): use embed pragma in future.
//
//go:generate go-bindata -pkg rulesdata -o rulesdata/rulesdata.go rules/rules.go

func init() {
	filename := "rules/rules.go"
	data, err := rulesdata.Asset(filename)
	if err != nil {
		panic(fmt.Sprintf("can't read embedded file: %v", err))
	}

	fset := token.NewFileSet()
	var groups []ruleguard.GoRuleGroup

	// First we create an Engine to parse all rules.
	// We need it to get the structured info about our rules
	// that will be used to generate checkers.
	// We introduce an extra scope in hope that rootEngine
	// will be garbage-collected after we don't need it.
	// LoadedGroups() returns a slice copy and that's all what we need.
	{
		rootEngine := ruleguard.NewEngine()

		parseContext := &ruleguard.ParseContext{
			Fset: fset,
		}
		if err := rootEngine.Load(parseContext, filename, bytes.NewReader(data)); err != nil {
			panic(fmt.Sprintf("load embedded ruleguard rules: %v", err))
		}
		groups = rootEngine.LoadedGroups()
	}

	// For every rules group we create a new checker and a separate engine.
	// That dedicated ruleguard engine will contain rules only from one group.
	for i := range groups {
		g := groups[i]
		info := &linter.CheckerInfo{
			Name:    g.Name,
			Summary: g.DocSummary,
			Before:  g.DocBefore,
			After:   g.DocAfter,
			Note:    g.DocNote,
			Tags:    g.DocTags,

			EmbeddedRuleguard: true,
		}
		collection.AddChecker(info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
			parseContext := &ruleguard.ParseContext{
				Fset: fset,
				GroupFilter: func(name string) bool {
					return name == g.Name
				},
			}
			engine := ruleguard.NewEngine()
			err := engine.Load(parseContext, filename, bytes.NewReader(data))
			if err != nil {
				return nil, err
			}
			c := &embeddedRuleguardChecker{
				ctx:    ctx,
				engine: engine,
			}
			return c, nil
		})
	}
}

type embeddedRuleguardChecker struct {
	ctx    *linter.CheckerContext
	engine *ruleguard.Engine
}

func (c *embeddedRuleguardChecker) WalkFile(f *ast.File) {
	runRuleguardEngine(c.ctx, f, c.engine, &ruleguard.RunContext{
		Pkg:   c.ctx.Pkg,
		Types: c.ctx.TypesInfo,
		Sizes: c.ctx.SizesInfo,
		Fset:  c.ctx.FileSet,
	})
}
