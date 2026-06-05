package compiler

import (
	"testing"

	"pricing-dsl/pkg/ast"
)

func TestBuildProgramCreatesDomainAST(t *testing.T) {
	tree, err := Parse(`
RULE "premium" IF user.type == "premium" AND cart.total > 1000 THEN percent 10 PRIORITY 20 GROUP coupon ENABLED
RULE "fallback" IF true THEN final 500
`)
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}

	program, err := BuildProgram(tree)
	if err != nil {
		t.Fatalf("build failed: %v", err)
	}

	if len(program.Rules) != 2 {
		t.Fatalf("expected 2 rules, got %d", len(program.Rules))
	}

	rule := program.Rules[0]
	if rule.Name != "premium" {
		t.Fatalf("expected rule name premium, got %s", rule.Name)
	}
	if rule.Priority == nil || *rule.Priority != 20 {
		t.Fatalf("expected priority 20, got %#v", rule.Priority)
	}
	if rule.Group == nil || *rule.Group != "coupon" {
		t.Fatalf("expected coupon group, got %#v", rule.Group)
	}
	if rule.Status == nil || *rule.Status != "ENABLED" {
		t.Fatalf("expected enabled status, got %#v", rule.Status)
	}
	if rule.Modifier.Type != ast.ModifierPercent || rule.Modifier.Value != 10 {
		t.Fatalf("unexpected modifier: %#v", rule.Modifier)
	}
	if _, ok := rule.Condition.(ast.LogicalExpression); !ok {
		t.Fatalf("expected logical expression, got %T", rule.Condition)
	}
	if _, ok := program.Rules[1].Condition.(ast.TrueExpression); !ok {
		t.Fatalf("expected true expression, got %T", program.Rules[1].Condition)
	}
}

func TestCompileReturnsValidationError(t *testing.T) {
	_, err := Compile(`RULE "bad" IF product.unknown == "x" THEN percent 10`)
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestCompileBuildsProgram(t *testing.T) {
	program, err := Compile(`RULE "premium" IF user.type == "premium" THEN fixed 100`)
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}
	if len(program.Rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(program.Rules))
	}
}

func TestParseReportsLexicalAndSyntacticErrors(t *testing.T) {
	tests := []string{
		`RULE "bad" IF user.type @ "premium" THEN fixed 100`,
		`RULE "bad" IF user.type == "premium" fixed 100`,
	}

	for _, source := range tests {
		if _, err := Parse(source); err == nil {
			t.Fatalf("expected parse error for %q", source)
		}
	}
}

func TestBuildProgramRejectsNilTree(t *testing.T) {
	_, err := BuildProgram(nil)
	if err == nil {
		t.Fatal("expected nil tree error")
	}
}
