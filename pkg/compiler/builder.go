package compiler

import (
	"fmt"
	"strconv"

	"github.com/kleo-53/pricing-dsl/parser"
	"github.com/kleo-53/pricing-dsl/pkg/ast"
)

func BuildProgram(tree parser.IStartContext) (*ast.Program, error) {
	if tree == nil {
		return nil, fmt.Errorf("start context is nil")
	}

	program := &ast.Program{}
	for _, ruleCtx := range tree.AllRule_() {
		rule, err := buildRule(ruleCtx)
		if err != nil {
			return nil, err
		}
		program.Rules = append(program.Rules, rule)
	}

	return program, nil
}

func buildRule(ctx parser.IRuleContext) (ast.Rule, error) {
	ruleCtx, ok := ctx.(*parser.RuleContext)
	if !ok {
		return ast.Rule{}, fmt.Errorf("unexpected rule context %T", ctx)
	}

	condition, err := buildExpression(ruleCtx.Expression())
	if err != nil {
		return ast.Rule{}, err
	}

	modifier, err := buildModifier(ruleCtx.Modifier())
	if err != nil {
		return ast.Rule{}, err
	}

	rule := ast.Rule{
		Name:      unquote(ruleCtx.STRING().GetText()),
		Condition: condition,
		Modifier:  modifier,
	}

	if ruleCtx.PRIORITY() != nil {
		priority, err := strconv.Atoi(ruleCtx.NUMBER().GetText())
		if err != nil {
			return ast.Rule{}, fmt.Errorf("invalid priority %q: %w", ruleCtx.NUMBER().GetText(), err)
		}
		rule.Priority = &priority
	}

	if ruleCtx.GROUP() != nil {
		group := ruleCtx.GroupType().GetText()
		rule.Group = &group
	}

	if ruleCtx.STATUS() != nil {
		status := ruleCtx.STATUS().GetText()
		rule.Status = &status
	}

	return rule, nil
}

func buildModifier(ctx parser.IModifierContext) (ast.Modifier, error) {
	modifierCtx, ok := ctx.(*parser.ModifierContext)
	if !ok {
		return ast.Modifier{}, fmt.Errorf("unexpected modifier context %T", ctx)
	}

	value, err := strconv.ParseFloat(modifierCtx.NUMBER().GetText(), 64)
	if err != nil {
		return ast.Modifier{}, fmt.Errorf("invalid modifier value %q: %w", modifierCtx.NUMBER().GetText(), err)
	}

	return ast.Modifier{
		Type:  ast.ModifierType(modifierCtx.ModifierType().GetText()),
		Value: value,
	}, nil
}

func buildExpression(ctx parser.IExpressionContext) (ast.Expression, error) {
	expressionCtx, ok := ctx.(*parser.ExpressionContext)
	if !ok {
		return nil, fmt.Errorf("unexpected expression context %T", ctx)
	}
	if expressionCtx.TRUE() != nil {
		return ast.TrueExpression{}, nil
	}

	comparisons := expressionCtx.AllComparison()
	if len(comparisons) == 0 {
		return nil, fmt.Errorf("expression has no comparisons")
	}

	left, err := buildComparison(comparisons[0])
	if err != nil {
		return nil, err
	}

	var expression ast.Expression = left
	for i := 1; i < len(comparisons); i++ {
		right, err := buildComparison(comparisons[i])
		if err != nil {
			return nil, err
		}

		operator := logicalOperatorAt(expressionCtx, i-1)
		expression = ast.LogicalExpression{
			Left:     expression,
			Operator: operator,
			Right:    right,
		}
	}

	return expression, nil
}

func buildComparison(ctx parser.IComparisonContext) (ast.Comparison, error) {
	comparisonCtx, ok := ctx.(*parser.ComparisonContext)
	if !ok {
		return ast.Comparison{}, fmt.Errorf("unexpected comparison context %T", ctx)
	}

	literal, err := buildLiteral(comparisonCtx.Value())
	if err != nil {
		return ast.Comparison{}, err
	}

	return ast.Comparison{
		Left:     buildIdentifier(comparisonCtx.Identifier()),
		Operator: ast.Comparator(comparisonCtx.Comparator().GetText()),
		Right:    literal,
	}, nil
}

func buildIdentifier(ctx parser.IIdentifierContext) ast.Identifier {
	identifierCtx := ctx.(*parser.IdentifierContext)
	parts := make([]string, 0, len(identifierCtx.AllID()))
	for _, token := range identifierCtx.AllID() {
		parts = append(parts, token.GetText())
	}

	return ast.Identifier{Parts: parts}
}

func buildLiteral(ctx parser.IValueContext) (ast.Literal, error) {
	valueCtx, ok := ctx.(*parser.ValueContext)
	if !ok {
		return ast.Literal{}, fmt.Errorf("unexpected value context %T", ctx)
	}

	if valueCtx.STRING() != nil {
		return ast.Literal{Type: ast.LiteralString, String: unquote(valueCtx.STRING().GetText())}, nil
	}

	value, err := strconv.ParseFloat(valueCtx.NUMBER().GetText(), 64)
	if err != nil {
		return ast.Literal{}, fmt.Errorf("invalid numeric literal %q: %w", valueCtx.NUMBER().GetText(), err)
	}

	return ast.Literal{Type: ast.LiteralNumber, Number: value}, nil
}

func logicalOperatorAt(ctx *parser.ExpressionContext, index int) ast.LogicalOperator {
	operatorIndex := 0
	for childIndex := 0; childIndex < ctx.GetChildCount(); childIndex++ {
		text := ctx.GetChild(childIndex).(interface{ GetText() string }).GetText()
		if text != string(ast.LogicalAnd) && text != string(ast.LogicalOr) {
			continue
		}
		if operatorIndex == index {
			return ast.LogicalOperator(text)
		}
		operatorIndex++
	}

	return ast.LogicalAnd
}

func unquote(value string) string {
	unquoted, err := strconv.Unquote(value)
	if err != nil {
		return value
	}

	return unquoted
}
