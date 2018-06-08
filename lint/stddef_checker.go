package lint

import (
	"go/ast"
	"go/constant"
	"go/parser"
	"go/token"
	"math"

	"github.com/Quasilyte/astcmp"
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

// TODO(quasilyte): should handle packages that define named constants
// specifically. For example, it's invalid to suggest math.* consts
// for math package itself.

// TODO(quasilyte): variables, types.
// For example: func(http.ResponseWriter, *http.Request) => http.HandlerFunc.

func stddefCheck(ctx *context) func(*ast.File) {
	expr := func(s string) ast.Expr {
		x, err := parser.ParseExpr(s)
		if err != nil {
			panic(err)
		}
		return x
	}

	return wrapLocalExprChecker(&stddefChecker{
		baseLocalExprChecker: baseLocalExprChecker{ctx: ctx},

		suggestionToExpression: map[string]ast.Expr{
			"math.MaxInt8":   expr(`1<<7 - 1`),
			"math.MinInt8":   expr(`-1 << 7`),
			"math.MaxInt16":  expr(`1<<15 - 1`),
			"math.MinInt16":  expr(`-1 << 15`),
			"math.MaxInt32":  expr(`1<<31 - 1`),
			"math.MinInt32":  expr(`-1 << 31`),
			"math.MaxInt64":  expr(`1<<63 - 1`),
			"math.MinInt64":  expr(`-1 << 63`),
			"math.MaxUint8":  expr(`1<<8 - 1`),
			"math.MaxUint16": expr(`1<<16 - 1`),
			"math.MaxUint32": expr(`1<<32 - 1`),
			"math.MaxUint64": expr(`1<<64 - 1`),
		},

		mathConsts: []mathConstant{
			{"math.Pi", math.Pi, 3.14},
			{"math.E", math.E, 2.71},

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
			"GET":    "net/http.MethodGet",
			"HEAD":   "net/http.MethodHead",
			"POST":   "net/http.MethodPost",
			"PUT":    "net/http.MethodPut",
			"DELETE": "net/http.MethodDelete",

			"Mon Jan _2 15:04:05 2006":            "time.ANSIC",
			"Mon Jan _2 15:04:05 MST 2006":        "time.UnixDate",
			"Mon Jan 02 15:04:05 -0700 2006":      "time.RubyDate",
			"02 Jan 06 15:04 MST":                 "time.RFC822",
			"02 Jan 06 15:04 -0700":               "time.RFC822Z",
			"Monday, 02-Jan-06 15:04:05 MST":      "time.RFC850",
			"Mon, 02 Jan 2006 15:04:05 MST":       "time.RFC1123",
			"Mon, 02 Jan 2006 15:04:05 -0700":     "time.RFC1123Z",
			"2006-01-02T15:04:05Z07:00":           "time.RFC3339",
			"2006-01-02T15:04:05.999999999Z07:00": "time.RFC3339Nano",

			"Jan _2 15:04:05":           "time.Stamp",
			"Jan _2 15:04:05.000":       "time.StampMilli",
			"Jan _2 15:04:05.000000":    "time.StampMicro",
			"Jan _2 15:04:05.000000000": "time.StampNano",

			"3:04PM": "time.Kitchen",
		},
	})
}

// mathConstant describes named constant value defined in "math" package.
type mathConstant struct {
	name  string
	value float64

	// imprecise is a common "short" value form.
	// Zero for constatns that don't have well-known short form.
	imprecise float64
}

type stddefChecker struct {
	// TODO(quasilyte): should be global expr checker. Refs #124.
	baseLocalExprChecker

	mathConsts []mathConstant

	stringLitToSuggestion map[string]string

	suggestionToExpression map[string]ast.Expr
}

func (c *stddefChecker) CheckLocalExpr(expr ast.Expr) {
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
			if astcmp.EqualExpr(expr, y) {
				c.warn(expr, suggestion)
				return
			}
		}
	case *ast.CallExpr:
		c.checkCallExpr(expr)
	}
}

func (c *stddefChecker) checkBasicLit(expr ast.Expr, val constant.Value) {
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

func (c *stddefChecker) checkCallExpr(call *ast.CallExpr) {
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

func (c *stddefChecker) warn(expr ast.Expr, suggestion string) {
	c.ctx.Warn(expr, "can replace %s with %s",
		nodeString(c.ctx.FileSet, expr), suggestion)
}
