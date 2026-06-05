package ast

import "testing"

func TestIdentifierString(t *testing.T) {
	tests := []struct {
		name       string
		identifier Identifier
		want       string
	}{
		{name: "empty", identifier: Identifier{}, want: ""},
		{name: "single", identifier: Identifier{Parts: []string{"cart"}}, want: "cart"},
		{name: "nested", identifier: Identifier{Parts: []string{"cart", "total"}}, want: "cart.total"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.identifier.String(); got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}
