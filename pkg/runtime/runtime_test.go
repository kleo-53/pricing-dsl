package runtime

import (
	"testing"

	"github.com/kleo-53/pricing-dsl/pkg/runtime/engine"
)

func TestNewReturnsEngine(t *testing.T) {
	var got interface{} = New()
	if _, ok := got.(engine.Engine); !ok {
		t.Fatalf("expected engine.Engine, got %T", got)
	}
}
