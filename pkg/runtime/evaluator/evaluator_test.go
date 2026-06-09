package evaluator

import (
	"testing"

	"github.com/kleo-53/pricing-dsl/pkg/ast"
	runtimecontext "github.com/kleo-53/pricing-dsl/pkg/runtime/context"
)

func TestEvaluateComparisonsAndLogicalExpressions(t *testing.T) {
	ctx := runtimecontext.Context{
		User: runtimecontext.UserContext{ID: "u-1", Type: "premium"},
		Product: runtimecontext.ProductContext{
			ID:       "p-1",
			Category: "electronics",
			Price:    1200,
		},
		Cart: runtimecontext.CartContext{Total: 1500},
	}

	tests := []struct {
		name       string
		expression ast.Expression
		want       bool
	}{
		{
			name: "string equality",
			expression: ast.Comparison{
				Left:     ast.Identifier{Parts: []string{"user", "type"}},
				Operator: ast.Equal,
				Right:    ast.Literal{Type: ast.LiteralString, String: "premium"},
			},
			want: true,
		},
		{
			name: "numeric greater than",
			expression: ast.Comparison{
				Left:     ast.Identifier{Parts: []string{"cart", "total"}},
				Operator: ast.GreaterThan,
				Right:    ast.Literal{Type: ast.LiteralNumber, Number: 1000},
			},
			want: true,
		},
		{
			name: "and expression",
			expression: ast.LogicalExpression{
				Left: ast.Comparison{
					Left:     ast.Identifier{Parts: []string{"product", "category"}},
					Operator: ast.Equal,
					Right:    ast.Literal{Type: ast.LiteralString, String: "electronics"},
				},
				Operator: ast.LogicalAnd,
				Right: ast.Comparison{
					Left:     ast.Identifier{Parts: []string{"product", "price"}},
					Operator: ast.LessThanOrEqual,
					Right:    ast.Literal{Type: ast.LiteralNumber, Number: 1200},
				},
			},
			want: true,
		},
		{
			name: "or expression",
			expression: ast.LogicalExpression{
				Left: ast.Comparison{
					Left:     ast.Identifier{Parts: []string{"user", "type"}},
					Operator: ast.Equal,
					Right:    ast.Literal{Type: ast.LiteralString, String: "guest"},
				},
				Operator: ast.LogicalOr,
				Right: ast.Comparison{
					Left:     ast.Identifier{Parts: []string{"cart", "total"}},
					Operator: ast.GreaterThanOrEqual,
					Right:    ast.Literal{Type: ast.LiteralNumber, Number: 1500},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expression, ctx)
			if err != nil {
				t.Fatalf("evaluate failed: %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestEvaluateReportsUnknownField(t *testing.T) {
	_, err := Evaluate(ast.Comparison{
		Left:     ast.Identifier{Parts: []string{"missing", "field"}},
		Operator: ast.Equal,
		Right:    ast.Literal{Type: ast.LiteralString, String: "x"},
	}, runtimecontext.Context{})
	if err == nil {
		t.Fatal("expected unknown field error")
	}
}

func TestEvaluateAdditionalBranches(t *testing.T) {
	ctx := runtimecontext.Context{
		User: runtimecontext.UserContext{Type: "premium", Age: 21},
		Cart: runtimecontext.CartContext{Total: 100, ItemsCount: 3},
	}

	tests := []struct {
		name       string
		expression ast.Expression
		want       bool
	}{
		{
			name: "not equal",
			expression: ast.Comparison{
				Left:     ast.Identifier{Parts: []string{"user", "type"}},
				Operator: ast.NotEqual,
				Right:    ast.Literal{Type: ast.LiteralString, String: "guest"},
			},
			want: true,
		},
		{
			name: "less than",
			expression: ast.Comparison{
				Left:     ast.Identifier{Parts: []string{"cart", "total"}},
				Operator: ast.LessThan,
				Right:    ast.Literal{Type: ast.LiteralNumber, Number: 200},
			},
			want: true,
		},
		{
			name: "legacy order field",
			expression: ast.Comparison{
				Left:     ast.Identifier{Parts: []string{"order", "total"}},
				Operator: ast.Equal,
				Right:    ast.Literal{Type: ast.LiteralNumber, Number: 100},
			},
			want: true,
		},
		{
			name: "legacy customer age field",
			expression: ast.Comparison{
				Left:     ast.Identifier{Parts: []string{"customer", "age"}},
				Operator: ast.GreaterThanOrEqual,
				Right:    ast.Literal{Type: ast.LiteralNumber, Number: 18},
			},
			want: true,
		},
		{
			name: "legacy items count field",
			expression: ast.Comparison{
				Left:     ast.Identifier{Parts: []string{"order", "itemsCount"}},
				Operator: ast.Equal,
				Right:    ast.Literal{Type: ast.LiteralNumber, Number: 3},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expression, ctx)
			if err != nil {
				t.Fatalf("evaluate failed: %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestEvaluateReportsUnsupportedExpression(t *testing.T) {
	_, err := Evaluate(nil, runtimecontext.Context{})
	if err == nil {
		t.Fatal("expected unsupported expression error")
	}
}

func TestEvaluateReportsNonNumericComparison(t *testing.T) {
	_, err := Evaluate(ast.Comparison{
		Left:     ast.Identifier{Parts: []string{"user", "type"}},
		Operator: ast.GreaterThan,
		Right:    ast.Literal{Type: ast.LiteralNumber, Number: 1},
	}, runtimecontext.Context{User: runtimecontext.UserContext{Type: "premium"}})
	if err == nil {
		t.Fatal("expected non-numeric comparison error")
	}
}
