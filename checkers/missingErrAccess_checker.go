package checkers

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"

	"github.com/go-critic/go-critic/framework/linter"
)

const recLimit = 20

var errRecLimit = fmt.Errorf(`recursion-limit-reached`)

func init() {
	var info linter.CheckerInfo
	info.Name = "missingErrAccess"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detect missing error access"
	info.Before = `
func processFileContent(path string) error {
	f,err := os.Open(path)	
	defer f.Close()
	if err != nil {
		return err
	}
	// ...
}`
	info.After = `
func processFileContent(path string) error {
	f,err := os.Open(path)	
	if err != nil {
		return err
	}
	defer f.Close()
	// ...
}
`
	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&missingErrAccess{ctx: ctx}), nil
	})
}

type missingErrAccess struct {
	ctx *linter.CheckerContext
	astwalk.WalkHandler
}

func (m *missingErrAccess) VisitFuncDecl(fnc *ast.FuncDecl) {
	if fnc.Body != nil {
		m.checkStmts(fnc.Body.List)
	}
}

func astSelectorExprToString(selector *ast.SelectorExpr) string {
	buf := &bytes.Buffer{}
	impAstSelectorExprToString(selector, buf, 0)
	return buf.String()
}

func impAstSelectorExprToString(selector *ast.SelectorExpr, buf *bytes.Buffer, recIdx int) {
	if recIdx >= recLimit {
		return
	}
	switch v := selector.X.(type) {
	case *ast.SelectorExpr:
		impAstSelectorExprToString(v, buf, recIdx+1)
	case *ast.Ident:
		buf.WriteString(v.Name)
	}
	if selector.Sel != nil {
		buf.WriteByte('.')
		buf.WriteString(selector.Sel.Name)
	}
}

func (m *missingErrAccess) checkStmts(stmts []ast.Stmt) {
	currErrOnLine := -1
	var currErrName string
	var currAssignStmt *ast.AssignStmt
	var currValues map[string]struct{}
	for _, stmt := range stmts {
		// check access
		if currErrOnLine != -1 {
			switch v := stmt.(type) {
			case *ast.RangeStmt:
				hasErrAcc, accessedValues, err := checkAccess(v.X, currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}
				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `range`, accessedValues)
				}

				if hasErrAcc || len(accessedValues) > 0 {
					currErrOnLine = -1
					break
				}
			case *ast.DeferStmt:

				hasErrAcc, accessedValues, err := checkAccess(v, currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}

				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `defer`, accessedValues)
				}

			case *ast.GoStmt:
				hasErrAcc, accessedValues, err := checkAccess(v.Call, currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}
				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `go`, accessedValues)
				}
			case *ast.IncDecStmt:
				hasErrAcc, accessedValues, err := checkAccess(v.X, currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}
				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `inc/dec`, accessedValues)
				}

				if hasErrAcc {
					currErrOnLine = -1
				}
			case *ast.AssignStmt:
				if len(v.Rhs) != 1 {
					break
				}

				hasErrAcc, accessedValues, err := checkAccess(v.Rhs[0], currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}
				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `assign`, accessedValues)
				}
				if hasErrAcc || len(accessedValues) > 0 {
					currErrOnLine = -1
					break
				}
			case *ast.ExprStmt:
				hasErrAcc, accessedValues, err := checkAccess(v, currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}
				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `expr`, accessedValues)
				}
				if hasErrAcc {
					currErrOnLine = -1
				}
			case *ast.DeclStmt:
				break
			case *ast.IfStmt:
				assign, ok := v.Init.(*ast.AssignStmt)
				if ok && len(assign.Rhs) == 1 {

					hasErrAcc, accessedValues, err := checkAccess(assign.Rhs[0], currValues, nil, 0)
					if err != nil {
						m.logRecFailed(stmt, `access`, err)
					}
					if !hasErrAcc && len(accessedValues) > 0 {
						m.log(v, `if`, accessedValues)
					}
				}

				hasErrAcc, accessedValues, err := checkAccess(v.Cond, currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}

				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `if(1)`, accessedValues)
				}

				currErrOnLine = -1
			case *ast.SendStmt:
				var shouldReset bool

				hasErrAcc, accessedValues, err := checkAccess(v.Value, currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}

				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `send`, accessedValues)
				}

				if hasErrAcc {
					currErrOnLine = -1
				}

				hasErrAcc, accessedValues, err = checkAccess(v.Chan, currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}

				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `send(1)`, accessedValues)
				}

				if hasErrAcc {
					currErrOnLine = -1
				}

				if shouldReset {
					currErrOnLine = -1
				}
			case *ast.TypeSwitchStmt, *ast.SwitchStmt:
				hasErrAcc, accessedValues, err := checkAccess(v, currValues, nil, 0)
				if err != nil {
					m.logRecFailed(stmt, `access`, err)
				}

				if !hasErrAcc && len(accessedValues) > 0 {
					m.log(v, `switch`, accessedValues)
				}

				if hasErrAcc || len(accessedValues) > 0 {
					currErrOnLine = -1
					break
				}
			case *ast.LabeledStmt,
				*ast.BranchStmt,
				*ast.SelectStmt,
				*ast.ReturnStmt:
				currErrOnLine = -1
			default:
				m.ctx.Warn(stmt, `statement not found`)
				currErrOnLine = -1
			}
		}

		topStmts, err := getTopStmts(stmt, nil, 0)
		if errors.Is(err, errRecLimit) {
			m.logRecFailed(stmt, `top-stmts`, err)
		} else if err != nil {
			m.ctx.Warn(stmt, fmt.Sprintf(`recursion failed %v`, err))
		}

		if err == nil {
			// check all statement slices
			for _, stmts := range topStmts {
				m.checkStmts(stmts)
			}
		}

		// check if values need to be reset
		nextErrOnLine := -1
		var nextAssignStmt *ast.AssignStmt
		var nextValues map[string]struct{}
		if assign, ok := stmt.(*ast.AssignStmt); ok {
			for _, exp := range assign.Lhs {
				if id, ok := exp.(*ast.Ident); ok {
					obj := m.ctx.TypesInfo.ObjectOf(id)
					if obj != nil {
						typeName := obj.Type().String()
						if typeName == `error` {
							if id.Name != `_` {
								nextErrOnLine = int(assign.TokPos)
								nextAssignStmt = assign
								currErrName = id.Name
							}
						}
					}
				}
			}

			if nextErrOnLine != -1 {
				nextValues = make(map[string]struct{})
				for _, exp := range assign.Lhs {
					switch v := exp.(type) {
					case *ast.Ident:
						if v.Name != currErrName {
							nextValues[v.Name] = struct{}{}
						}

					case *ast.SelectorExpr:
						nextValues[astSelectorExprToString(v)] = struct{}{}
					}
				}
			}
		}

		if nextErrOnLine != -1 {
			if currErrOnLine != -1 {
				m.ctx.Warn(currAssignStmt, `no error access`)
			}

			currErrOnLine = nextErrOnLine
			currAssignStmt = nextAssignStmt
			currValues = nextValues
		}
	}
}

func (m missingErrAccess) log(stmt ast.Stmt, msg string, accessedValues []string) {
	m.ctx.Warn(stmt, fmt.Sprintf(`%v, missing error check accessing %v`, msg, accessedValues))
}

func (m missingErrAccess) logRecFailed(stmt ast.Stmt, msg string, err error) {
	m.ctx.Warn(stmt, fmt.Sprintf(`%v recursion failed %v`, msg, err))
}

func checkAccess(node interface{},
	values map[string]struct{},
	currAccessedValues []string,
	recIdx int) (hasErrAccess bool, accessedValues []string, err error) {

	if recIdx >= recLimit {
		return false, accessedValues, errRecLimit
	}

	switch v := node.(type) {
	case *ast.SelectorExpr:
		sel := astSelectorExprToString(v)

		if strings.HasPrefix(sel, `err.`) {
			return true, currAccessedValues, nil
		}

		for {
			_, ok := values[sel]
			if ok {
				currAccessedValues = append(currAccessedValues, sel)
				break
			}

			idx := strings.LastIndex(sel, `.`)
			if idx == -1 {
				break
			}

			sel = sel[:idx]
		}
		return false, currAccessedValues, nil
	case *ast.BlockStmt:
		return false, currAccessedValues, nil
	case *ast.Ident:
		switch v.Name {
		case `err`:
			hasErrAccess = true
		default:
			if _, ok := values[v.Name]; ok {
				currAccessedValues = append(currAccessedValues, v.Name)
			}
		}

		return hasErrAccess, currAccessedValues, nil
	default:
		val := reflect.ValueOf(v)

		switch val.Kind() {
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				childVal := val.Index(i)
				if childVal.CanInterface() {
					var errAccess bool
					var err error
					errAccess, currAccessedValues, err = checkAccess(childVal.Interface(), values, currAccessedValues, recIdx+1)
					if err != nil {
						return hasErrAccess, currAccessedValues, err

					}

					if errAccess {
						hasErrAccess = true
					}
				}
			}
			return hasErrAccess, currAccessedValues, nil
		case reflect.Ptr:
			val := val.Elem()
			if val.Kind() == reflect.Struct {
				var errAccess bool
				var err error
				errAccess, currAccessedValues, err = checkAccessStructValue(val, values, currAccessedValues, recIdx)
				if err != nil {
					return hasErrAccess, currAccessedValues, err
				}

				if errAccess {
					hasErrAccess = true
				}
			}
			return hasErrAccess, currAccessedValues, nil
		case reflect.Struct:
			var errAccess bool
			var err error
			errAccess, currAccessedValues, err = checkAccessStructValue(val, values, currAccessedValues, recIdx)
			if err != nil {
				return hasErrAccess, currAccessedValues, err
			}

			if errAccess {
				hasErrAccess = true
			}
			return hasErrAccess, currAccessedValues, nil
		default:
			return false, currAccessedValues, nil
		}
	}
}

func checkAccessStructValue(val reflect.Value,
	values map[string]struct{},
	currAccessedValues []string,
	recIdx int) (hasErrAccess bool, accessedValues []string, err error) {

	for i := 0; i < val.NumField(); i++ {
		childVal := val.Field(i)
		if childVal.CanInterface() {
			var errAccess bool
			var err error
			errAccess, currAccessedValues, err = checkAccess(childVal.Interface(), values, currAccessedValues, recIdx+1)
			if err != nil {
				return hasErrAccess, currAccessedValues, err
			}

			if errAccess {
				hasErrAccess = true
			}
		}
	}

	return hasErrAccess, currAccessedValues, nil
}

func getTopStmts(node interface{}, stmts [][]ast.Stmt, recIdx int) ([][]ast.Stmt, error) {

	if recIdx >= recLimit {
		return stmts, errRecLimit
	}

	switch v := node.(type) {
	case *ast.BlockStmt:
		return append(stmts, v.List), nil
	case []ast.Stmt:
		return append(stmts, v), nil
	case *ast.Object:
		return stmts, nil
	default:
		val := reflect.ValueOf(v)
		switch val.Kind() {
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				childVal := val.Index(i)
				if childVal.CanInterface() {
					var err error
					stmts, err = getTopStmts(childVal.Interface(), stmts, recIdx+1)
					if err != nil {
						return stmts, err
					}
				}
			}
			return stmts, nil
		case reflect.Ptr:
			val := val.Elem()
			if val.Kind() == reflect.Struct {
				var err error
				stmts, err = getTopStmtsWithStruct(val, stmts, recIdx)
				if err != nil {
					return stmts, err
				}
			}
			return stmts, nil
		case reflect.Struct:
			var err error
			stmts, err = getTopStmtsWithStruct(val, stmts, recIdx)
			return stmts, err
		default:
			return stmts, nil
		}
	}
}

func getTopStmtsWithStruct(val reflect.Value, stmts [][]ast.Stmt, recIdx int) ([][]ast.Stmt, error) {
	var err error
	for i := 0; i < val.NumField(); i++ {
		childVal := val.Field(i)
		if childVal.CanInterface() {
			stmts, err = getTopStmts(childVal.Interface(), stmts, recIdx+1)
		}
	}

	return stmts, err
}
