package compiler

import (
	"fmt"
	"strings"

	"pricing-dsl/lexer"
	"pricing-dsl/parser"
	"pricing-dsl/pkg/ast"
	"pricing-dsl/validator"

	"github.com/antlr4-go/antlr/v4"
)

func Compile(source string) (*ast.Program, error) {
	result := validator.ValidateString(source)
	if result.Stage != validator.StageValid {
		return nil, fmt.Errorf("%s errors: %s", result.Stage, strings.Join(result.Errors, "; "))
	}

	tree, err := Parse(source)
	if err != nil {
		return nil, err
	}

	return BuildProgram(tree)
}

func Parse(source string) (parser.IStartContext, error) {
	input := antlr.NewInputStream(source)
	collector := &lexer.ErrorCollector{}

	lex := parser.NewPricingDSLLexer(input)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(collector)

	stream := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	stream.Fill()
	if len(collector.LexerErrors) > 0 {
		return nil, fmt.Errorf("lexical errors: %s", strings.Join(collector.LexerErrors, "; "))
	}

	stream.Seek(0)
	p := parser.NewPricingDSLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(collector)

	tree := p.Start_()
	if len(collector.ParserErrors) > 0 {
		return nil, fmt.Errorf("syntactic errors: %s", strings.Join(collector.ParserErrors, "; "))
	}

	return tree, nil
}
