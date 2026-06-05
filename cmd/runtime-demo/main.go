package main

import (
	"fmt"
	"os"

	"pricing-dsl/pkg/compiler"
	"pricing-dsl/pkg/runtime"
	runtimecontext "pricing-dsl/pkg/runtime/context"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("provide a DSL rules file to execute")
		os.Exit(1)
	}

	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("failed to read rules file:", err)
		os.Exit(1)
	}

	program, err := compiler.Compile(string(source))
	if err != nil {
		fmt.Println("failed to compile DSL:", err)
		os.Exit(1)
	}

	engine := runtime.New()
	result, err := engine.Execute(program, runtimecontext.Context{
		User: runtimecontext.UserContext{
			ID:   "u-1",
			Type: "premium",
		},
		Product: runtimecontext.ProductContext{
			ID:       "p-1",
			Category: "electronics",
			Price:    1200,
		},
		Cart: runtimecontext.CartContext{
			Total: 1200,
		},
	})
	if err != nil {
		fmt.Println("failed to execute rules:", err)
		os.Exit(1)
	}

	fmt.Printf("Original price: %.2f\n", result.OriginalPrice)
	fmt.Printf("Final price: %.2f\n", result.FinalPrice)
	fmt.Println("Applied rules:")
	for _, rule := range result.AppliedRules {
		fmt.Printf("- %s: %s %.2f\n", rule.Name, rule.ModifierType, rule.Value)
	}
}
