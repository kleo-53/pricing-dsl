package ast

type Program struct {
	Rules []Rule
}

type Rule struct {
	Name      string
	Condition Expression
	Modifier  Modifier
	Priority  *int
	Group     *string
	Status    *string
}

type Modifier struct {
	Type  ModifierType
	Value float64
}

type ModifierType string

const (
	ModifierPercent ModifierType = "percent"
	ModifierFixed   ModifierType = "fixed"
	ModifierFinal   ModifierType = "final"
)

type Expression interface {
	expressionNode()
}

type TrueExpression struct{}

func (TrueExpression) expressionNode() {}

type Comparison struct {
	Left     Identifier
	Operator Comparator
	Right    Literal
}

func (Comparison) expressionNode() {}

type LogicalExpression struct {
	Left     Expression
	Operator LogicalOperator
	Right    Expression
}

func (LogicalExpression) expressionNode() {}

type LogicalOperator string

const (
	LogicalAnd LogicalOperator = "AND"
	LogicalOr  LogicalOperator = "OR"
)

type Comparator string

const (
	Equal              Comparator = "=="
	NotEqual           Comparator = "!="
	GreaterThanOrEqual Comparator = ">="
	GreaterThan        Comparator = ">"
	LessThanOrEqual    Comparator = "<="
	LessThan           Comparator = "<"
)

type Identifier struct {
	Parts []string
}

func (i Identifier) String() string {
	if len(i.Parts) == 0 {
		return ""
	}

	result := i.Parts[0]
	for _, part := range i.Parts[1:] {
		result += "." + part
	}

	return result
}

type Literal struct {
	Type   LiteralType
	String string
	Number float64
}

type LiteralType string

const (
	LiteralString LiteralType = "string"
	LiteralNumber LiteralType = "number"
)
