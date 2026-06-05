package model

type AppliedRule struct {
	Name         string
	ModifierType string
	Value        float64
}

type ExecutionResult struct {
	OriginalPrice float64
	FinalPrice    float64
	AppliedRules  []AppliedRule
}
