package validator

import (
	"strings"
	"testing"
)

const validProgram = `
RULE "adult-percent" IF customer.age >= 18 THEN percent 10 PRIORITY 10 GROUP coupon
RULE "segment-fixed" IF customer.segment == "vip" AND order.total > 1000 THEN fixed 250
RULE "fallback" IF true THEN final 0 GROUP cert DISABLED
`

func TestValidateStringAcceptsValidProgram(t *testing.T) {
	result := ValidateString(validProgram)
	if result.Stage != StageValid {
		t.Fatalf("expected valid program, got %s errors: %v", result.Stage, result.Errors)
	}
}

func TestValidateStringReportsExactlyOneFailureAtEachAnalysisStage(t *testing.T) {
	tests := []struct {
		name          string
		source        string
		expectedStage Stage
		expectedText  []string
	}{
		{
			name:          "lexical",
			source:        `RULE "bad-token" IF customer.age @ 18 THEN percent 10`,
			expectedStage: StageLexical,
			expectedText: []string{"line 1:33 token recognition error at: '@'"},
		},
		{
			name:          "syntactic",
			source:        `RULE "missing-then" IF customer.age > 18 percent 10`,
			expectedStage: StageSyntactic,
			expectedText: []string{"line 1:41 missing 'THEN' at 'percent"},
		},
		{
			name: "semantic",
			source: `
RULE "duplicate" IF customer.age > 18 THEN percent 150 PRIORITY 101
RULE "duplicate" IF customer.unknown == "vip" THEN fixed 100001
`,
			expectedStage: StageSemantic,
			expectedText: []string{
				"percent modifier 150.00 outside range [-100..100]",
				"priority 101 outside range [0..100]",
				"duplicate rule name \"duplicate\"",
				"unknown field customer.unknown",
				"fixed modifier 100001.00 is too large",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateString(tt.source)
			if result.Stage != tt.expectedStage {
				t.Fatalf("expected %s stage, got %s errors: %v", tt.expectedStage, result.Stage, result.Errors)
			}
			if len(result.Errors) == 0 {
				t.Fatalf("expected at least one %s error", tt.expectedStage)
			}
			messages := strings.Join(result.Errors, "\n")
			for _, want := range tt.expectedText {
				if !strings.Contains(messages, want) {
					t.Fatalf("expected %s error %q in:\n%s", tt.expectedStage, want, messages)
				}
			}
		})
	}
}
