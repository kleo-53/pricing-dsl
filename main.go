package main

import (
	"fmt"
	"os"

	"pricing-dsl/validator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("provide a DSL rules file to analyze")
		os.Exit(1)
	}

	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("failed to read rules file:", err)
		os.Exit(1)
	}

	result := validator.ValidateString(string(source))
	if result.Stage == validator.StageValid {
		fmt.Println("DSL is valid")
		return
	}

	fmt.Printf("%s errors:\n", result.Stage)
	for _, err := range result.Errors {
		fmt.Println("-", err)
	}
	os.Exit(1)
}
