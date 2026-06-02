// Code generated from grammar/PricingDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PricingDSL
import "github.com/antlr4-go/antlr/v4"

type BasePricingDSLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePricingDSLVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePricingDSLVisitor) VisitRule(ctx *RuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePricingDSLVisitor) VisitModifier(ctx *ModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePricingDSLVisitor) VisitModifierType(ctx *ModifierTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePricingDSLVisitor) VisitGroupType(ctx *GroupTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePricingDSLVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePricingDSLVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePricingDSLVisitor) VisitComparator(ctx *ComparatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePricingDSLVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePricingDSLVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
