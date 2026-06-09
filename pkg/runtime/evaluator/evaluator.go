package evaluator

import (
	"fmt"
	"strconv"

	"github.com/kleo-53/pricing-dsl/pkg/ast"
	runtimecontext "github.com/kleo-53/pricing-dsl/pkg/runtime/context"
)

func Evaluate(expression ast.Expression, ctx runtimecontext.Context) (bool, error) {
	switch expression := expression.(type) {
	case ast.TrueExpression:
		return true, nil
	case ast.Comparison:
		return evaluateComparison(expression, ctx)
	case ast.LogicalExpression:
		left, err := Evaluate(expression.Left, ctx)
		if err != nil {
			return false, err
		}

		switch expression.Operator {
		case ast.LogicalAnd:
			if !left {
				return false, nil
			}
			return Evaluate(expression.Right, ctx)
		case ast.LogicalOr:
			if left {
				return true, nil
			}
			return Evaluate(expression.Right, ctx)
		default:
			return false, fmt.Errorf("unsupported logical operator %s", expression.Operator)
		}
	default:
		return false, fmt.Errorf("unsupported expression %T", expression)
	}
}

func evaluateComparison(comparison ast.Comparison, ctx runtimecontext.Context) (bool, error) {
	left, err := resolveIdentifier(comparison.Left, ctx)
	if err != nil {
		return false, err
	}

	right := literalValue(comparison.Right)

	switch comparison.Operator {
	case ast.Equal:
		return compareEqual(left, right), nil
	case ast.NotEqual:
		return !compareEqual(left, right), nil
	case ast.GreaterThan, ast.GreaterThanOrEqual, ast.LessThan, ast.LessThanOrEqual:
		leftNumber, rightNumber, err := numericPair(left, right)
		if err != nil {
			return false, err
		}
		switch comparison.Operator {
		case ast.GreaterThan:
			return leftNumber > rightNumber, nil
		case ast.GreaterThanOrEqual:
			return leftNumber >= rightNumber, nil
		case ast.LessThan:
			return leftNumber < rightNumber, nil
		case ast.LessThanOrEqual:
			return leftNumber <= rightNumber, nil
		}
	}

	return false, fmt.Errorf("unsupported comparator %s", comparison.Operator)
}

func resolveIdentifier(identifier ast.Identifier, ctx runtimecontext.Context) (interface{}, error) {
	switch identifier.String() {
	case "user.id":
		return ctx.User.ID, nil
	case "user.type":
		return ctx.User.Type, nil
	case "product.id":
		return ctx.Product.ID, nil
	case "product.category":
		return ctx.Product.Category, nil
	case "product.price":
		return ctx.Product.Price, nil
	case "cart.total":
		return ctx.Cart.Total, nil
	case "customer.age":
		return ctx.User.Age, nil
	case "customer.segment":
		return ctx.User.Type, nil
	case "order.total":
		return ctx.Cart.Total, nil
	case "order.itemsCount":
		return ctx.Cart.ItemsCount, nil
	default:
		return nil, fmt.Errorf("unknown field %s", identifier.String())
	}
}

func literalValue(literal ast.Literal) interface{} {
	if literal.Type == ast.LiteralNumber {
		return literal.Number
	}

	return literal.String
}

func compareEqual(left, right interface{}) bool {
	leftNumber, rightNumber, err := numericPair(left, right)
	if err == nil {
		return leftNumber == rightNumber
	}

	return fmt.Sprint(left) == fmt.Sprint(right)
}

func numericPair(left, right interface{}) (float64, float64, error) {
	leftNumber, err := toFloat(left)
	if err != nil {
		return 0, 0, err
	}
	rightNumber, err := toFloat(right)
	if err != nil {
		return 0, 0, err
	}

	return leftNumber, rightNumber, nil
}

func toFloat(value interface{}) (float64, error) {
	switch value := value.(type) {
	case float64:
		return value, nil
	case int:
		return float64(value), nil
	case string:
		return strconv.ParseFloat(value, 64)
	default:
		return 0, fmt.Errorf("value %v is not numeric", value)
	}
}
