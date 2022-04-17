package ssautil

import (
	"context"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ast/inspector"
)

type contextKey int

const (
	_ contextKey = iota
	passKey
	ssaKey
	inspectorKey
	posMapKey
)

func Pass(ctx context.Context) *analysis.Pass {
	return ctx.Value(passKey).(*analysis.Pass)
}

func SSA(ctx context.Context) *buildssa.SSA {
	return ctx.Value(ssaKey).(*buildssa.SSA)
}

func Inspector(ctx context.Context) *inspector.Inspector {
	return ctx.Value(inspectorKey).(*inspector.Inspector)
}

func PosMap(ctx context.Context) map[token.Pos][]ast.Node {
	return ctx.Value(posMapKey).(map[token.Pos][]ast.Node)
}
