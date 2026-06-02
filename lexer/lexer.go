package lexer

import (
	"fmt"
	"pricing-dsl/parser"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorCollector struct {
	*antlr.DefaultErrorListener

	LexerErrors  []string
	ParserErrors []string
}

func (e *ErrorCollector) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	ex antlr.RecognitionException,
) {

	errMsg := fmt.Sprintf(
		"строка %d:%d %s",
		line,
		column,
		msg,
	)

	switch recognizer.(type) {

	case *parser.PricingDSLLexer:
		e.LexerErrors = append(e.LexerErrors, errMsg)

	case *parser.PricingDSLParser:
		e.ParserErrors = append(e.ParserErrors, errMsg)

	default:
		e.ParserErrors = append(e.ParserErrors, errMsg)
	}
}