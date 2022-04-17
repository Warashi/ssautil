package ssautil

import (
	"go/ast"
	"go/token"
	"strconv"
)

func ReplaceConstArg(expr *ast.CallExpr, actual, want string) *ast.CallExpr {
	replaced := *expr
	replaced.Args = make([]ast.Expr, len(expr.Args))
	copy(replaced.Args, expr.Args)

	for i, arg := range replaced.Args {
		c, ok := arg.(*ast.BasicLit)
		if !ok || c.Kind != token.STRING {
			continue
		}
		if strconv.Quote(actual) == c.Value || (strconv.CanBackquote(actual) && c.Value == "`"+actual+"`") {
			replaced.Args[i] = &ast.BasicLit{
				ValuePos: c.ValuePos,
				Kind:     c.Kind,
				Value:    strconv.Quote(want),
			}
			return &replaced
		}
	}

	return &replaced
}
