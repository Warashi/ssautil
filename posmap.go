package ssautil

import (
	"context"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

type Poser interface {
	Pos() token.Pos
}

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

func Node[U ast.Node, T Poser](ctx context.Context, p T) (node U, ok bool) {
	posMap := PosMap(ctx)
	for i := p.Pos(); i > 0; i-- {
		stack := posMap[i]
		if len(stack) == 0 {
			break
		}
		for j := range stack {
			node := stack[len(stack)-1-j]
			ident, ok := node.(U)
			if ok {
				return ident, true
			}
		}
	}
	var zero U
	return zero, false
}
