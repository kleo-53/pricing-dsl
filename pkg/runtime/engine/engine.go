package engine

import (
	"sort"

	"github.com/kleo-53/pricing-dsl/pkg/ast"
	"github.com/kleo-53/pricing-dsl/pkg/model"
	runtimecontext "github.com/kleo-53/pricing-dsl/pkg/runtime/context"
	"github.com/kleo-53/pricing-dsl/pkg/runtime/evaluator"
)

type Context = runtimecontext.Context
type AppliedRule = model.AppliedRule
type ExecutionResult = model.ExecutionResult

type Engine interface {
	Execute(program *ast.Program, ctx Context) (ExecutionResult, error)
}

type defaultEngine struct{}

func New() Engine {
	return defaultEngine{}
}

func (e defaultEngine) Execute(program *ast.Program, ctx Context) (ExecutionResult, error) {
	result := ExecutionResult{
		OriginalPrice: ctx.Product.Price,
		FinalPrice:    ctx.Product.Price,
	}
	if program == nil {
		return result, nil
	}

	rules := append([]ast.Rule(nil), program.Rules...)
	sort.SliceStable(rules, func(i, j int) bool {
		return priorityOf(rules[i]) > priorityOf(rules[j])
	})

	for _, rule := range rules {
		if rule.Status != nil && *rule.Status == "DISABLED" {
			continue
		}

		matches, err := evaluator.Evaluate(rule.Condition, ctx)
		if err != nil {
			return result, err
		}
		if !matches {
			continue
		}

		result.FinalPrice = applyModifier(result.FinalPrice, rule.Modifier)
		result.AppliedRules = append(result.AppliedRules, AppliedRule{
			Name:         rule.Name,
			ModifierType: string(rule.Modifier.Type),
			Value:        rule.Modifier.Value,
		})
	}

	return result, nil
}

func priorityOf(rule ast.Rule) int {
	if rule.Priority == nil {
		return 0
	}

	return *rule.Priority
}

func applyModifier(price float64, modifier ast.Modifier) float64 {
	switch modifier.Type {
	case ast.ModifierPercent:
		return price * (1 - modifier.Value/100)
	case ast.ModifierFixed:
		return price - modifier.Value
	case ast.ModifierFinal:
		return modifier.Value
	default:
		return price
	}
}
