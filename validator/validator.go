package validator

import (
	"fmt"
	"strconv"

	"github.com/kleo-53/pricing-dsl/lexer"
	"github.com/kleo-53/pricing-dsl/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Stage string

const (
	StageValid     Stage = "valid"
	StageLexical   Stage = "lexical"
	StageSyntactic Stage = "syntactic"
	StageSemantic  Stage = "semantic"
)

type Result struct {
	Stage  Stage
	Errors []string
}

func ValidateString(source string) Result {
	input := antlr.NewInputStream(source)
	collector := &lexer.ErrorCollector{}

	lex := parser.NewPricingDSLLexer(input)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(collector)

	stream := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	stream.Fill()
	if len(collector.LexerErrors) > 0 {
		return Result{Stage: StageLexical, Errors: collector.LexerErrors}
	}

	stream.Seek(0)
	p := parser.NewPricingDSLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(collector)

	tree := p.Start_()
	if len(collector.ParserErrors) > 0 {
		return Result{Stage: StageSyntactic, Errors: collector.ParserErrors}
	}

	v := NewValidator()
	tree.Accept(v)
	if len(v.Errors()) > 0 {
		return Result{Stage: StageSemantic, Errors: v.Errors()}
	}

	return Result{Stage: StageValid}
}

type Validator struct {
	parser.BasePricingDSLVisitor

	errors    []string
	ruleNames map[string]bool
}

func NewValidator() *Validator {
	return &Validator{
		errors:    []string{},
		ruleNames: make(map[string]bool),
	}
}

func (v *Validator) Errors() []string {
	return v.errors
}

func (v *Validator) addError(msg string) {
	v.errors = append(v.errors, msg)
}

func (v *Validator) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Validator) VisitChildren(node antlr.RuleNode) interface{} {
	for i := 0; i < node.GetChildCount(); i++ {
		if child, ok := node.GetChild(i).(antlr.ParseTree); ok {
			child.Accept(v)
		}
	}

	return nil
}

func (v *Validator) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (v *Validator) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

func (v *Validator) VisitStart(ctx *parser.StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitComparison(ctx *parser.ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Validator) VisitRule(ctx *parser.RuleContext) interface{} {
	ruleName := ctx.STRING().GetText()

	if v.ruleNames[ruleName] {
		v.addError(fmt.Sprintf("duplicate rule name %s", ruleName))
	}
	v.ruleNames[ruleName] = true

	if ctx.PRIORITY() != nil {
		numberToken := ctx.NUMBER()
		priority, err := strconv.Atoi(numberToken.GetText())
		if err != nil {
			v.addError(fmt.Sprintf("invalid priority in rule %s", ruleName))
		} else if priority < 0 || priority > 100 {
			v.addError(
				fmt.Sprintf("priority %d outside range [0..100] in rule %s", priority, ruleName),
			)
		}
	}

	return v.VisitChildren(ctx)
}

var allowedIdentifiers = map[string]bool{
	"customer.age":     true,
	"customer.segment": true,
	"order.total":      true,
	"order.itemsCount": true,
	"user.id":          true,
	"user.type":        true,
	"product.id":       true,
	"product.category": true,
	"product.price":    true,
	"cart.total":       true,
}

func (v *Validator) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	id := ctx.GetText()
	if !allowedIdentifiers[id] {
		v.addError(fmt.Sprintf("unknown field %s", id))
	}

	return nil
}

func (v *Validator) VisitModifier(ctx *parser.ModifierContext) interface{} {
	value, err := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
	if err != nil {
		return nil
	}

	switch {
	case ctx.ModifierType().PERCENT() != nil:
		if value < -100 || value > 100 {
			v.addError(fmt.Sprintf("percent modifier %.2f outside range [-100..100]", value))
		}
	case ctx.ModifierType().FIXED() != nil:
		if value > 100000 {
			v.addError(fmt.Sprintf("fixed modifier %.2f is too large", value))
		}
	}

	return nil
}
