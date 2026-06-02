// Code generated from grammar/PricingDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PricingDSL
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PricingDSLParser.
type PricingDSLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PricingDSLParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by PricingDSLParser#rule.
	VisitRule(ctx *RuleContext) interface{}

	// Visit a parse tree produced by PricingDSLParser#modifier.
	VisitModifier(ctx *ModifierContext) interface{}

	// Visit a parse tree produced by PricingDSLParser#modifierType.
	VisitModifierType(ctx *ModifierTypeContext) interface{}

	// Visit a parse tree produced by PricingDSLParser#groupType.
	VisitGroupType(ctx *GroupTypeContext) interface{}

	// Visit a parse tree produced by PricingDSLParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by PricingDSLParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by PricingDSLParser#comparator.
	VisitComparator(ctx *ComparatorContext) interface{}

	// Visit a parse tree produced by PricingDSLParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by PricingDSLParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
