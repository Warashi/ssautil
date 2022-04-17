package ssautil

import (
	"context"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/ssa"
)

func Prepare(ctx context.Context, pass *analysis.Pass) context.Context {
	ctx = context.WithValue(ctx, passKey, pass)
	ctx = context.WithValue(ctx, ssaKey, pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA))
	ctx = context.WithValue(ctx, inspectorKey, pass.ResultOf[inspect.Analyzer].(*inspector.Inspector))
	ctx = context.WithValue(ctx, posMapKey, buildPosMap(ctx))
	return ctx
}

func buildPosMap(ctx context.Context) map[token.Pos][]ast.Node {
	posMap := make(map[token.Pos][]ast.Node)
	Inspector(ctx).Preorder(nil, func(node ast.Node) {
		for i := node.Pos(); i <= node.End(); i++ {
			posMap[i] = append(posMap[i], node)
		}
	})
	return posMap
}

func CallExpr(ctx context.Context, call *ssa.Call) (*ast.CallExpr, bool) {
	posMap := PosMap(ctx)
	for i := call.Pos(); i > 0; i-- {
		stack := posMap[i]
		if len(stack) == 0 {
			break
		}
		for j := range stack {
			node := stack[len(stack)-1-j]
			ident, ok := node.(*ast.CallExpr)
			if ok {
				return ident, true
			}
		}
	}
	return nil, false
}
