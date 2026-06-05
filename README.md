# Pricing DSL

Pricing DSL is a small Go platform for parsing, validating, compiling, versioning,
publishing, and executing pricing rules for store-like integrations.

## Architecture

```text
grammar/                 ANTLR grammar source
parser/                  generated ANTLR lexer, parser, visitor, listener
lexer/                   shared ANTLR error collector
validator/               lexical, syntactic, and semantic validation
pkg/ast/                 hand-written domain AST with no ANTLR dependencies
pkg/compiler/            parser facade and AST builder
pkg/model/               runtime result models
pkg/runtime/             SDK runtime facade
pkg/runtime/context/     store integration context
pkg/runtime/evaluator/   AST expression evaluator
pkg/runtime/engine/      rule execution engine
pkg/repository/          rule version repository abstraction
pkg/release/             publication and rollback manager
cmd/runtime-demo/        CLI demo for compile and execute
```

## DSL Example

```text
RULE "premium-percent" IF user.type == "premium" AND cart.total > 1000 THEN percent 10 PRIORITY 20 GROUP coupon
RULE "electronics-fixed" IF product.category == "electronics" THEN fixed 100 PRIORITY 10
RULE "fallback-final" IF true THEN final 500 DISABLED
```

## SDK Usage

```go
program, err := compiler.Compile(dsl)
if err != nil {
	return err
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
```

The SDK is intentionally independent from HTTP, databases, message brokers, and
UI frameworks.

## Runtime Semantics

Rules are executed in descending priority order. Missing priority is treated as
`0`. Disabled rules are skipped.

Modifiers are applied to the current price:

```text
percent 10  => price = price * 0.9
fixed 100   => price = price - 100
final 500   => price = 500
```

## Versioning And Release

`pkg/repository` defines the rule-version storage contract and an in-memory
implementation. `pkg/release` manages active versions, deactivation, and
rollback. No database or broker integration is included yet.

## Demo

```text
go run ./cmd/runtime-demo ./rules.dsl
```

## Tests

```text
go test ./...
go test ./pkg/... -cover
```
