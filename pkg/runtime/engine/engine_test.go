package engine

import (
	"testing"

	"pricing-dsl/pkg/ast"
	runtimecontext "pricing-dsl/pkg/runtime/context"
)

func TestExecuteAppliesModifiersInPriorityOrder(t *testing.T) {
	priorityHigh := 20
	priorityLow := 10
	program := &ast.Program{Rules: []ast.Rule{
		{
			Name:      "fixed-second",
			Condition: ast.TrueExpression{},
			Modifier:  ast.Modifier{Type: ast.ModifierFixed, Value: 100},
			Priority:  &priorityLow,
		},
		{
			Name:      "percent-first",
			Condition: ast.TrueExpression{},
			Modifier:  ast.Modifier{Type: ast.ModifierPercent, Value: 10},
			Priority:  &priorityHigh,
		},
	}}

	result, err := New().Execute(program, runtimecontext.Context{
		Product: runtimecontext.ProductContext{Price: 1000},
	})
	if err != nil {
		t.Fatalf("execute failed: %v", err)
	}

	if result.FinalPrice != 800 {
		t.Fatalf("expected final price 800, got %.2f", result.FinalPrice)
	}
	if len(result.AppliedRules) != 2 {
		t.Fatalf("expected 2 applied rules, got %d", len(result.AppliedRules))
	}
	if result.AppliedRules[0].Name != "percent-first" {
		t.Fatalf("expected percent-first to apply first, got %s", result.AppliedRules[0].Name)
	}
}

func TestExecuteSupportsFinalModifierAndDisabledRules(t *testing.T) {
	disabled := "DISABLED"
	program := &ast.Program{Rules: []ast.Rule{
		{
			Name:      "disabled",
			Condition: ast.TrueExpression{},
			Modifier:  ast.Modifier{Type: ast.ModifierFixed, Value: 1000},
			Status:    &disabled,
		},
		{
			Name:      "final",
			Condition: ast.TrueExpression{},
			Modifier:  ast.Modifier{Type: ast.ModifierFinal, Value: 500},
		},
	}}

	result, err := New().Execute(program, runtimecontext.Context{
		Product: runtimecontext.ProductContext{Price: 1200},
	})
	if err != nil {
		t.Fatalf("execute failed: %v", err)
	}

	if result.OriginalPrice != 1200 || result.FinalPrice != 500 {
		t.Fatalf("unexpected prices: %#v", result)
	}
	if len(result.AppliedRules) != 1 || result.AppliedRules[0].Name != "final" {
		t.Fatalf("unexpected applied rules: %#v", result.AppliedRules)
	}
}

func TestExecuteAppliesMatchingRuleOnly(t *testing.T) {
	program := &ast.Program{Rules: []ast.Rule{
		{
			Name: "premium",
			Condition: ast.Comparison{
				Left:     ast.Identifier{Parts: []string{"user", "type"}},
				Operator: ast.Equal,
				Right:    ast.Literal{Type: ast.LiteralString, String: "premium"},
			},
			Modifier: ast.Modifier{Type: ast.ModifierPercent, Value: 10},
		},
	}}

	result, err := New().Execute(program, runtimecontext.Context{
		User:    runtimecontext.UserContext{Type: "guest"},
		Product: runtimecontext.ProductContext{Price: 1000},
	})
	if err != nil {
		t.Fatalf("execute failed: %v", err)
	}
	if result.FinalPrice != 1000 || len(result.AppliedRules) != 0 {
		t.Fatalf("unexpected result: %#v", result)
	}
}
