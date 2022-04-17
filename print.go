package ssautil

import (
	"bytes"
	"context"
	"go/ast"
	"go/format"
)

func PrettyPrint(ctx context.Context, expr ast.Expr) []byte {
	pass := Pass(ctx)
	var b bytes.Buffer
	format.Node(&b, pass.Fset, expr)
	return b.Bytes()
}
