package release

import (
	"testing"

	"github.com/kleo-53/pricing-dsl/pkg/repository"
)

func TestManagerPublishesAndRollsBack(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	for _, version := range []string{"v1", "v2"} {
		if err := repo.Save(repository.RuleVersion{Version: version}); err != nil {
			t.Fatalf("save %s failed: %v", version, err)
		}
	}

	manager := NewManager(repo)
	if err := manager.Publish("v1"); err != nil {
		t.Fatalf("publish v1 failed: %v", err)
	}
	if err := manager.Publish("v2"); err != nil {
		t.Fatalf("publish v2 failed: %v", err)
	}
	if manager.ActiveVersion() != "v2" {
		t.Fatalf("expected active v2, got %s", manager.ActiveVersion())
	}

	if err := manager.Rollback(""); err != nil {
		t.Fatalf("rollback failed: %v", err)
	}
	if manager.ActiveVersion() != "v1" {
		t.Fatalf("expected active v1, got %s", manager.ActiveVersion())
	}
}

func TestManagerDeactivate(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	if err := repo.Save(repository.RuleVersion{Version: "v1"}); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	manager := NewManager(repo)
	if err := manager.Publish("v1"); err != nil {
		t.Fatalf("publish failed: %v", err)
	}
	if err := manager.Deactivate("v1"); err != nil {
		t.Fatalf("deactivate failed: %v", err)
	}
	if manager.ActiveVersion() != "" {
		t.Fatalf("expected no active version, got %s", manager.ActiveVersion())
	}
}
