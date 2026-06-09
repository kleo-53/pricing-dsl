package release

import (
	"fmt"

	"github.com/kleo-53/pricing-dsl/pkg/repository"
)

type Manager interface {
	Publish(version string) error
	Deactivate(version string) error
	Rollback(version string) error
	ActiveVersion() string
}

type manager struct {
	repository repository.Repository
	history    []string
}

func NewManager(repository repository.Repository) Manager {
	return &manager{repository: repository}
}

func (m *manager) Publish(version string) error {
	if version == "" {
		return fmt.Errorf("version is required")
	}

	active := m.ActiveVersion()
	if active != "" && active != version {
		m.history = append(m.history, active)
	}

	return m.repository.SetActiveVersion(version)
}

func (m *manager) Deactivate(version string) error {
	active := m.ActiveVersion()
	if active != version {
		return fmt.Errorf("version %s is not active", version)
	}

	return m.repository.DeactivateActiveVersion()
}

func (m *manager) Rollback(version string) error {
	if version == "" {
		if len(m.history) == 0 {
			return fmt.Errorf("no previous version")
		}
		version = m.history[len(m.history)-1]
		m.history = m.history[:len(m.history)-1]
	}

	return m.repository.SetActiveVersion(version)
}

func (m *manager) ActiveVersion() string {
	active, err := m.repository.GetActiveVersion()
	if err != nil {
		return ""
	}

	return active.Version
}
