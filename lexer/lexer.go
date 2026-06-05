package lexer

import (
	"fmt"

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
	errMsg := fmt.Sprintf("line %d:%d %s", line, column, msg)

	if offendingSymbol == nil {
		e.LexerErrors = append(e.LexerErrors, errMsg)
		return
	}

	e.ParserErrors = append(e.ParserErrors, errMsg)
}
