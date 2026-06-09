package repository

import (
	"fmt"
	"sort"
	"time"

	"github.com/kleo-53/pricing-dsl/pkg/ast"
)

type RuleVersion struct {
	ID        string
	Version   string
	CreatedAt time.Time
	Program   *ast.Program
}

type Repository interface {
	Save(version RuleVersion) error
	Get(version string) (*RuleVersion, error)
	ListVersions() ([]RuleVersion, error)
	GetActiveVersion() (*RuleVersion, error)
	SetActiveVersion(version string) error
	DeactivateActiveVersion() error
}

type InMemoryRepository struct {
	versions      map[string]RuleVersion
	activeVersion string
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{versions: make(map[string]RuleVersion)}
}

func (r *InMemoryRepository) Save(version RuleVersion) error {
	if version.Version == "" {
		return fmt.Errorf("version is required")
	}
	if version.ID == "" {
		version.ID = version.Version
	}
	if version.CreatedAt.IsZero() {
		version.CreatedAt = time.Now()
	}

	r.versions[version.Version] = version
	return nil
}

func (r *InMemoryRepository) Get(version string) (*RuleVersion, error) {
	ruleVersion, ok := r.versions[version]
	if !ok {
		return nil, fmt.Errorf("version %s not found", version)
	}

	return &ruleVersion, nil
}

func (r *InMemoryRepository) ListVersions() ([]RuleVersion, error) {
	versions := make([]RuleVersion, 0, len(r.versions))
	for _, version := range r.versions {
		versions = append(versions, version)
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].CreatedAt.Before(versions[j].CreatedAt)
	})

	return versions, nil
}

func (r *InMemoryRepository) GetActiveVersion() (*RuleVersion, error) {
	if r.activeVersion == "" {
		return nil, fmt.Errorf("no active version")
	}

	return r.Get(r.activeVersion)
}

func (r *InMemoryRepository) SetActiveVersion(version string) error {
	if _, err := r.Get(version); err != nil {
		return err
	}

	r.activeVersion = version
	return nil
}

func (r *InMemoryRepository) DeactivateActiveVersion() error {
	r.activeVersion = ""
	return nil
}
