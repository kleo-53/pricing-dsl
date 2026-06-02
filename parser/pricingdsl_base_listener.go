// Code generated from grammar/PricingDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PricingDSL
import "github.com/antlr4-go/antlr/v4"

// BasePricingDSLListener is a complete listener for a parse tree produced by PricingDSLParser.
type BasePricingDSLListener struct{}

var _ PricingDSLListener = &BasePricingDSLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePricingDSLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePricingDSLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePricingDSLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePricingDSLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BasePricingDSLListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BasePricingDSLListener) ExitStart(ctx *StartContext) {}

// EnterRule is called when production rule is entered.
func (s *BasePricingDSLListener) EnterRule(ctx *RuleContext) {}

// ExitRule is called when production rule is exited.
func (s *BasePricingDSLListener) ExitRule(ctx *RuleContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BasePricingDSLListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BasePricingDSLListener) ExitModifier(ctx *ModifierContext) {}

// EnterModifierType is called when production modifierType is entered.
func (s *BasePricingDSLListener) EnterModifierType(ctx *ModifierTypeContext) {}

// ExitModifierType is called when production modifierType is exited.
func (s *BasePricingDSLListener) ExitModifierType(ctx *ModifierTypeContext) {}

// EnterGroupType is called when production groupType is entered.
func (s *BasePricingDSLListener) EnterGroupType(ctx *GroupTypeContext) {}

// ExitGroupType is called when production groupType is exited.
func (s *BasePricingDSLListener) ExitGroupType(ctx *GroupTypeContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePricingDSLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePricingDSLListener) ExitExpression(ctx *ExpressionContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BasePricingDSLListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BasePricingDSLListener) ExitComparison(ctx *ComparisonContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BasePricingDSLListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BasePricingDSLListener) ExitComparator(ctx *ComparatorContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasePricingDSLListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasePricingDSLListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterValue is called when production value is entered.
func (s *BasePricingDSLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasePricingDSLListener) ExitValue(ctx *ValueContext) {}
