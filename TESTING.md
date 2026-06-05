# Testing Rules

Validator tests are staged:

1. Lexical analysis tested with invalid input that contains a token the lexer cannot recognize.
2. Syntactic analysis tested with invalid input made only from valid tokens arranged against the grammar.
3. Semantic analysis tested with syntactically valid invalid input; bundle semantic rule violations into that case when adding coverage.
4. Valid programs should continue to cover accepted grammar shapes and must not depend on error text.

The validator must stop at the first failing stage, so lexical failures do not also assert parser or semantic behavior.
