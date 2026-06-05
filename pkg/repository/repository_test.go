package repository

import (
	"testing"
	"time"

	"pricing-dsl/pkg/ast"
)

func TestInMemoryRepositoryStoresVersionsAndActiveVersion(t *testing.T) {
	repo := NewInMemoryRepository()
	if err := repo.Save(RuleVersion{
		Version:   "v1",
		CreatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		Program:   &ast.Program{},
	}); err != nil {
		t.Fatalf("save v1 failed: %v", err)
	}
	if err := repo.Save(RuleVersion{
		Version:   "v2",
		CreatedAt: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
		Program:   &ast.Program{},
	}); err != nil {
		t.Fatalf("save v2 failed: %v", err)
	}

	versions, err := repo.ListVersions()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(versions) != 2 || versions[0].Version != "v1" || versions[1].Version != "v2" {
		t.Fatalf("unexpected versions: %#v", versions)
	}

	if err := repo.SetActiveVersion("v2"); err != nil {
		t.Fatalf("set active failed: %v", err)
	}
	active, err := repo.GetActiveVersion()
	if err != nil {
		t.Fatalf("get active failed: %v", err)
	}
	if active.Version != "v2" {
		t.Fatalf("expected active v2, got %s", active.Version)
	}
}

func TestInMemoryRepositoryReportsErrors(t *testing.T) {
	repo := NewInMemoryRepository()

	if err := repo.Save(RuleVersion{}); err == nil {
		t.Fatal("expected save without version to fail")
	}
	if _, err := repo.Get("missing"); err == nil {
		t.Fatal("expected missing version to fail")
	}
	if _, err := repo.GetActiveVersion(); err == nil {
		t.Fatal("expected missing active version to fail")
	}
	if err := repo.SetActiveVersion("missing"); err == nil {
		t.Fatal("expected set active missing version to fail")
	}
	if err := repo.DeactivateActiveVersion(); err != nil {
		t.Fatalf("deactivate without active should be harmless: %v", err)
	}
}
