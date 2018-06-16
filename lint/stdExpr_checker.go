package lint

import (
	"go/ast"
	"go/constant"
	"go/token"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/go-toolsmith/astequal"
	"github.com/go-toolsmith/strparse"
)

// TODO(quasilyte): if we were tracking expression context, it would be possible
// to do more suggestions with less false positive rate.
//
// Right now we're somewhat conservative in some cases, and very eager in
// others. For example, "UPDATE" literal can be unrelated to HTTP, but
// checker still emits a warning (context awareness could help here).
//
// In numerical comparison cases we could suggest size constants instead of
// literals like 0x7fffffff => math.MaxInt32, because bitwise masks
// are not used in this way.

// TODO(quasilyte): variables, types.
// For example: func(http.ResponseWriter, *http.Request) => http.HandlerFunc.

func init() {
	addChecker(stdExprChecker{}, &ruleInfo{})
}

// mathConstant describes named constant value defined in "math" package.
type mathConstant struct {
	name  string
	value float64

	// imprecise is a common "short" value form.
	// Zero for constatns that don't have well-known short form.
	imprecise float64
}

type stdExprChecker struct {
	baseExprChecker

	mathConsts []mathConstant

	stringLitToSuggestion map[string]string

	suggestionToExpression map[string]ast.Expr
}

func (c stdExprChecker) New(ctx *context) func(*ast.File) {
	return wrapExprChecker(&stdExprChecker{
		baseExprChecker: baseExprChecker{ctx: ctx},

		suggestionToExpression: map[string]ast.Expr{
			"math.MaxInt8":   strparse.Expr(`1<<7 - 1`),
			"math.MinInt8":   strparse.Expr(`-1 << 7`),
			"math.MaxInt16":  strparse.Expr(`1<<15 - 1`),
			"math.MinInt16":  strparse.Expr(`-1 << 15`),
			"math.MaxInt32":  strparse.Expr(`1<<31 - 1`),
			"math.MinInt32":  strparse.Expr(`-1 << 31`),
			"math.MaxInt64":  strparse.Expr(`1<<63 - 1`),
			"math.MinInt64":  strparse.Expr(`-1 << 63`),
			"math.MaxUint8":  strparse.Expr(`1<<8 - 1`),
			"math.MaxUint16": strparse.Expr(`1<<16 - 1`),
			"math.MaxUint32": strparse.Expr(`1<<32 - 1`),
			"math.MaxUint64": strparse.Expr(`1<<64 - 1`),
		},

		mathConsts: []mathConstant{
			// Unary plus is a current way to avoid stddef to trigger
			// on these literals.
			{"math.Pi", math.Pi, +3.14},
			{"math.E", math.E, +2.71},

			{"math.Phi", math.Phi, 0},

			{"math.Sqrt2", math.Sqrt2, 0},
			{"math.SqrtE", math.SqrtE, 0},
			{"math.SqrtPi", math.SqrtPi, 0},
			{"math.SqrtPhi", math.SqrtPhi, 0},

			{"math.Ln2", math.Ln2, 0},
			{"math.Log2E", math.Log2E, 0},
			{"math.Ln10", math.Ln10, 0},
			{"math.Log10E", math.Log10E, 0},
		},

		stringLitToSuggestion: map[string]string{
			http.MethodGet:    "net/http.MethodGet",
			http.MethodHead:   "net/http.MethodHead",
			http.MethodPost:   "net/http.MethodPost",
			http.MethodPut:    "net/http.MethodPut",
			http.MethodDelete: "net/http.MethodDelete",

			time.ANSIC:       "time.ANSIC",
			time.UnixDate:    "time.UnixDate",
			time.RubyDate:    "time.RubyDate",
			time.RFC822:      "time.RFC822",
			time.RFC822Z:     "time.RFC822Z",
			time.RFC850:      "time.RFC850",
			time.RFC1123:     "time.RFC1123",
			time.RFC1123Z:    "time.RFC1123Z",
			time.RFC3339:     "time.RFC3339",
			time.RFC3339Nano: "time.RFC3339Nano",
			time.Stamp:       "time.Stamp",
			time.StampMilli:  "time.StampMilli",
			time.StampMicro:  "time.StampMicro",
			time.StampNano:   "time.StampNano",
			time.Kitchen:     "time.Kitchen",
		},
	})
}

func (c *stdExprChecker) CheckExpr(expr ast.Expr) {
	val := c.ctx.TypesInfo.Types[expr].Value
	if val == nil {
		// Not a compile-time constant.
		return
	}

	switch expr := expr.(type) {
	case *ast.BasicLit:
		c.checkBasicLit(expr, val)
	case *ast.BinaryExpr:
		for suggestion, y := range c.suggestionToExpression {
			if astequal.Expr(expr, y) {
				c.warn(expr, suggestion)
				return
			}
		}
	case *ast.CallExpr:
		c.checkCallExpr(expr)
	}
}

func (c *stdExprChecker) checkBasicLit(expr ast.Expr, val constant.Value) {
	const epsilon = 0.00001

	switch val.Kind() {
	case constant.Float:
		v, ok := constant.Float64Val(val)
		if !ok {
			return
		}
		for _, mathConst := range c.mathConsts {
			cond := math.Abs(v-mathConst.value) < epsilon ||
				(mathConst.imprecise != 0 && mathConst.imprecise == v)
			if cond {
				c.warn(expr, mathConst.name)
				return
			}
		}

	case constant.String:
		// TODO(quasilyte): it's better to check these only in a proper context
		// when we know that we're dealing with HTTP-related code.
		// We could also suggest HTTP status code constants in that
		// context as it will reduce false positive change to a minimum.
		// If we can infer time-related context, then we can suggest
		// non-string time package constants (like time.Second).
		suggestion := c.stringLitToSuggestion[constant.StringVal(val)]
		if suggestion != "" {
			c.warn(expr, suggestion)
			return
		}
	}
}

func (c *stdExprChecker) checkCallExpr(call *ast.CallExpr) {
	if qualifiedName(call.Fun) == "unsafe.Sizeof" {
		switch x := call.Args[0].(type) {
		case *ast.BasicLit:
			if x.Kind == token.INT {
				c.warn(call, "math/bits.UintSize")
			}
		case *ast.CallExpr:
			fn, ok := x.Fun.(*ast.Ident)
			if !ok {
				return
			}
			// TODO: is there a constant for "uintptr" size?
			switch fn.Name {
			case "uint", "int":
				c.warn(call, "math/bits.UintSize")
			}
		}
	}
}

func (c *stdExprChecker) warn(expr ast.Expr, suggestion string) {
	// Avoid printing warnings for packages that use recognized
	// expressions to define constants/variables we are suggesting.
	definingPkg := strings.Split(suggestion, ".")[0]
	if c.ctx.Package.Name() != definingPkg {
		c.ctx.Warn(expr, "can replace %s with %s", expr, suggestion)
	}
}
