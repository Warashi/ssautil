package ssautil

import (
	"golang.org/x/exp/slices"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var requires = []*analysis.Analyzer{
	inspect.Analyzer,
	buildssa.Analyzer,
}

func Requires(a ...*analysis.Analyzer) []*analysis.Analyzer {
	for _, r := range requires {
		if !slices.Contains(a, r) {
			a = append(a, r)
		}
	}
	return a
}
