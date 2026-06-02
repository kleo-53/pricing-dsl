package validator

import (
	"fmt"
	"strconv"

	"pricing-dsl/parser"
)

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

func (v *Validator) VisitRule(ctx *parser.RuleContext) interface{} {

	fmt.Println("VISIT RULE")
	// RULE "name"
	ruleName := ctx.STRING().GetText()

	// Проверка уникальности имени
	if v.ruleNames[ruleName] {
		v.addError(
			fmt.Sprintf("дублирующееся имя правила %s", ruleName),
		)
	}

	v.ruleNames[ruleName] = true

	// Проверка PRIORITY
	if ctx.PRIORITY() != nil {

		// NUMBER после PRIORITY
		numberToken := ctx.NUMBER()
		priority, err := strconv.Atoi(numberToken.GetText())

		if err != nil {
			v.addError(
				fmt.Sprintf("неверный приоритет в правиле %s", ruleName),
			)
		} else if priority < 0 || priority > 100 {
			v.addError(
				fmt.Sprintf(
					"приоритет %d вне диапазона [0..100] в правиле %s",
					priority,
					ruleName,
				),
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
}

func (v *Validator) VisitIdentifier(
	ctx *parser.IdentifierContext,
) interface{} {

	id := ctx.GetText()

	if !allowedIdentifiers[id] {
		v.addError(
			fmt.Sprintf("неизвестное поле %s", id),
		)
	}

	return nil
}

func (v *Validator) VisitModifier(
	ctx *parser.ModifierContext,
) interface{} {

	value, err := strconv.ParseFloat(
		ctx.NUMBER().GetText(),
		64,
	)

	if err != nil {
		return nil
	}

	switch {

	case ctx.ModifierType().PERCENT() != nil:
		if value < -100 || value > 100 {
			v.addError(
				fmt.Sprintf(
					"percent modifier %.2f вне диапазона [-100..100]",
					value,
				),
			)
		}

	case ctx.ModifierType().FIXED() != nil:
		if value > 100000 {
			v.addError(
				fmt.Sprintf(
					"слишком большой fixed modifier %.2f",
					value,
				),
			)
		}
	}

	return nil
}
