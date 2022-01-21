package checkers

import (
	"fmt"
	"go/ast"
	"strings"
	"testing"
)

type mockLogger struct {
	logs []string
}

func (m *mockLogger) Warn(node ast.Node, format string, args ...interface{}) {
	m.logs = append(m.logs, fmt.Sprintf(format, args...))
}

func TestMissingErrAccess_errRecLimit(t *testing.T) {

	assign := &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.Ident{
				Name: `err`,
			},
			&ast.Ident{
				Name: `val`,
			},
		},
	}

	// create recursion loop
	elt := &ast.Ellipsis{}
	elt.Elt = elt

	expr := &ast.ExprStmt{X: elt}

	stmts := []ast.Stmt{
		assign,
		expr,
	}

	mlog := &mockLogger{}
	ec := &missingErrAccess{logger: mlog}

	ec.checkStmts(stmts)

	if len(mlog.logs) != 2 {
		t.FailNow()
	}

	msg := strings.Join(mlog.logs, `,`)
	if !strings.Contains(msg, `recursion-limit-reached`) ||
		!strings.Contains(msg, `access`) ||
		!strings.Contains(msg, `top-stmts`) {
		t.Fail()
	}
}
