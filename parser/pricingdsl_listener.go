// Code generated from grammar/PricingDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PricingDSL
import "github.com/antlr4-go/antlr/v4"

// PricingDSLListener is a complete listener for a parse tree produced by PricingDSLParser.
type PricingDSLListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterRule is called when entering the rule production.
	EnterRule(c *RuleContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterModifierType is called when entering the modifierType production.
	EnterModifierType(c *ModifierTypeContext)

	// EnterGroupType is called when entering the groupType production.
	EnterGroupType(c *GroupTypeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitRule is called when exiting the rule production.
	ExitRule(c *RuleContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitModifierType is called when exiting the modifierType production.
	ExitModifierType(c *ModifierTypeContext)

	// ExitGroupType is called when exiting the groupType production.
	ExitGroupType(c *GroupTypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
