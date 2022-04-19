package ssautil

import (
	"go/types"

	"golang.org/x/tools/go/ssa"
)

type Referrerer interface {
	Referrers() *[]ssa.Instruction
}

type Operander interface {
	Operands([]*ssa.Value) []*ssa.Value
}

func Referrers(r Referrerer) []ssa.Instruction {
	refs := r.Referrers()
	if refs == nil {
		return nil
	}
	nonnil := make([]ssa.Instruction, 0, len(*refs))
	for i := range *refs {
		if (*refs)[i] != nil {
			nonnil = append(nonnil, (*refs)[i])
		}
	}
	return nonnil
}

func Operands(o Operander) []ssa.Value {
	ops := o.Operands(nil)
	nonnil := make([]ssa.Value, 0, len(ops))
	for i := range ops {
		if ops[i] != nil && *ops[i] != nil {
			nonnil = append(nonnil, *ops[i])
		}
	}
	return nonnil
}

func IsContext(v *types.Var) bool {
	return v.Type().String() == "context.Context"
}
