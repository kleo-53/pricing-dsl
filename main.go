package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"pricing-dsl/lexer"
	"pricing-dsl/parser"
	"pricing-dsl/validator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите файл с правилами для анализа")
		os.Exit(1)
	}

	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	collector := &lexer.ErrorCollector{}

	lex := parser.NewPricingDSLLexer(input)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(collector)

	stream := antlr.NewCommonTokenStream(
		lex,
		antlr.TokenDefaultChannel,
	)
	if len(collector.LexerErrors) > 0 {
		fmt.Println("Лексические ошибки:")
		for _, err := range collector.LexerErrors {
			fmt.Println("-", err)
		}
		return
	}

	p := parser.NewPricingDSLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(collector)
	tree := p.Start_()

	if len(collector.ParserErrors) > 0 {
		fmt.Println("Синтаксические ошибки:")
		for _, err := range collector.ParserErrors {
			fmt.Println("-", err)
		}
		return
	}
	v := validator.NewValidator()
	tree.Accept(v)

	fmt.Println("DSL корректен")
}
