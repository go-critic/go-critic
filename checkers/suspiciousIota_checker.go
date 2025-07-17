package checkers

import (
	"go/ast"
	"go/constant"
	"go/token"
	"strings"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"

	"github.com/go-toolsmith/astfmt"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "suspiciousIota"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects suspicious iota usage patterns"
	info.Before = `
const (
	A = iota
	B = iota // redundant iota
	C = iota // redundant iota
)`
	info.After = `
const (
	A = iota
	B        // auto-increment
	C        // auto-increment
)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return &suspiciousIotaChecker{ctx: ctx}, nil
	})
}

type suspiciousIotaChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *suspiciousIotaChecker) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.CONST {
			c.checkConstBlock(genDecl)
		}
	}
}

type iotaInfo struct {
	hasIota      bool
	hasExplicit  bool
	iotaExprs    []ast.Expr
	explicitVals []ast.Expr
	specs        []*ast.ValueSpec
}

func (c *suspiciousIotaChecker) checkConstBlock(decl *ast.GenDecl) {
	if len(decl.Specs) <= 1 {
		return
	}

	info := c.analyzeConstBlock(decl)

	c.checkRedundantIota(info)
	c.checkInconsistentIota(info)
	c.checkMixedIotaExplicit(info)
}

func (c *suspiciousIotaChecker) analyzeConstBlock(decl *ast.GenDecl) *iotaInfo {
	info := &iotaInfo{
		iotaExprs:    make([]ast.Expr, 0),
		explicitVals: make([]ast.Expr, 0),
		specs:        make([]*ast.ValueSpec, 0),
	}

	for _, spec := range decl.Specs {
		if valueSpec, ok := spec.(*ast.ValueSpec); ok {
			info.specs = append(info.specs, valueSpec)

			for _, value := range valueSpec.Values {
				if c.containsIota(value) {
					info.hasIota = true
					info.iotaExprs = append(info.iotaExprs, value)
				} else if value != nil {
					info.hasExplicit = true
					info.explicitVals = append(info.explicitVals, value)
				}
			}
		}
	}

	return info
}

func (c *suspiciousIotaChecker) containsIota(expr ast.Expr) bool {
	found := false
	ast.Inspect(expr, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Name == "iota" {
			found = true
			return false
		}
		return true
	})
	return found
}

func (c *suspiciousIotaChecker) checkRedundantIota(info *iotaInfo) {
	if !info.hasIota || len(info.specs) < 2 {
		return
	}

	firstHasIota := false
	for i, spec := range info.specs {
		hasExplicitIota := false
		for _, value := range spec.Values {
			if c.isExplicitIota(value) {
				hasExplicitIota = true
				break
			}
		}

		if i == 0 {
			firstHasIota = hasExplicitIota
		} else if hasExplicitIota && firstHasIota {
			for _, value := range spec.Values {
				if c.isExplicitIota(value) {
					c.warnRedundantIota(value, spec.Names[0])
				}
			}
		}
	}
}

func (c *suspiciousIotaChecker) isExplicitIota(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name == "iota"
	}
	return false
}

func (c *suspiciousIotaChecker) checkInconsistentIota(info *iotaInfo) {
	if !info.hasIota || len(info.iotaExprs) < 2 {
		return
	}

	exprStrings := make(map[string]int)
	var firstExpr string
	firstSpecIdx := -1

	for i, spec := range info.specs {
		if len(spec.Values) > 0 && c.containsIota(spec.Values[0]) {
			firstExpr = astfmt.Sprint(spec.Values[0])
			firstSpecIdx = i
			break
		}
	}

	if firstSpecIdx == -1 {
		return
	}

	for i, spec := range info.specs {
		if i <= firstSpecIdx || len(spec.Values) == 0 {
			continue
		}

		if spec.Values[0] != nil {
			exprStr := astfmt.Sprint(spec.Values[0])
			if c.containsIota(spec.Values[0]) {
				exprStrings[exprStr]++
			}
		}
	}

	if len(exprStrings) > 1 {
		for i, spec := range info.specs {
			if i <= firstSpecIdx || len(spec.Values) == 0 || spec.Values[0] == nil {
				continue
			}

			if c.containsIota(spec.Values[0]) {
				exprStr := astfmt.Sprint(spec.Values[0])
				if exprStr != firstExpr && !c.isExplicitIota(spec.Values[0]) {
					c.warnInconsistentIota(spec.Values[0], spec.Names[0], firstExpr)
				}
			}
		}
	}
}

func (c *suspiciousIotaChecker) checkMixedIotaExplicit(info *iotaInfo) {
	if !info.hasIota || !info.hasExplicit {
		return
	}

	if c.looksLikeBitFlags(info) || c.looksLikeAcceptablePattern(info) {
		return
	}

	firstIotaIdx := -1
	for i, spec := range info.specs {
		for _, value := range spec.Values {
			if c.containsIota(value) {
				firstIotaIdx = i
				break
			}
		}
		if firstIotaIdx != -1 {
			break
		}
	}

	for i, spec := range info.specs {
		for j, value := range spec.Values {
			if value != nil && !c.containsIota(value) && c.hasConstantValue(value) {
				var name *ast.Ident
				if j < len(spec.Names) {
					name = spec.Names[j]
				} else if len(spec.Names) > 0 {
					name = spec.Names[0]
				}
				if name != nil {
					if firstIotaIdx != -1 && i < firstIotaIdx {
						c.warnExplicitBeforeIota(value, name)
					} else {
						c.warnMixedIotaExplicit(value, name)
					}
				}
			}
		}
	}
}

func (c *suspiciousIotaChecker) looksLikeBitFlags(info *iotaInfo) bool {
	for _, spec := range info.specs {
		for _, value := range spec.Values {
			if value != nil {
				valueStr := astfmt.Sprint(value)
				if valueStr == "0" || strings.Contains(valueStr, "<<") || strings.Contains(valueStr, "1 <<") {
					return true
				}
			}
		}
	}
	return false
}

func (c *suspiciousIotaChecker) looksLikeAcceptablePattern(info *iotaInfo) bool {
	explicitCount := 0
	lastExplicitIsEnd := false

	for i, spec := range info.specs {
		for _, value := range spec.Values {
			if value != nil && !c.containsIota(value) && c.hasConstantValue(value) {
				explicitCount++
				if i == len(info.specs)-1 {
					for _, name := range spec.Names {
						nameStr := strings.ToLower(name.Name)
						if strings.Contains(nameStr, "end") || strings.Contains(nameStr, "max") ||
							strings.Contains(nameStr, "last") || strings.Contains(nameStr, "none") {
							lastExplicitIsEnd = true
						}
					}
				}
			}
		}
	}

	return explicitCount == 1 && lastExplicitIsEnd
}

func (c *suspiciousIotaChecker) hasConstantValue(expr ast.Expr) bool {
	switch expr.(type) {
	case *ast.BasicLit:
		return true
	case *ast.Ident:
		return true
	}

	if c.ctx.TypesInfo != nil {
		if tv, ok := c.ctx.TypesInfo.Types[expr]; ok && tv.Value != nil {
			return tv.Value.Kind() != constant.Unknown
		}
	}

	return false
}

func (c *suspiciousIotaChecker) warnRedundantIota(expr ast.Expr, name *ast.Ident) {
	c.ctx.Warn(expr, "redundant iota usage for %s; iota auto-increments without explicit assignment", name.Name)
}

func (c *suspiciousIotaChecker) warnInconsistentIota(expr ast.Expr, name *ast.Ident, expected string) {
	current := astfmt.Sprint(expr)
	c.ctx.Warn(expr, "inconsistent iota pattern for %s: got %s, expected %s", name.Name, current, expected)
}

func (c *suspiciousIotaChecker) warnMixedIotaExplicit(expr ast.Expr, name *ast.Ident) {
	c.ctx.Warn(expr, "mixing explicit values with iota in const block may be confusing for %s", name.Name)
}

func (c *suspiciousIotaChecker) warnExplicitBeforeIota(expr ast.Expr, name *ast.Ident) {
	c.ctx.Warn(expr, "const %s appears before iota usage; this affects iota values and may cause bugs", name.Name)
}
