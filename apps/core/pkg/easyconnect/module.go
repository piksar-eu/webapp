package easyconnect

import (
	"github.com/piksar-eu/webapp/apps/core/pkg/auth"
)

const PermissionLeadsView = "leads_view"

func LoadModule(leadRepo LeadRepository, authStore auth.AuthorizationStore) *Module {
	m := &Module{
		leadRepo:  leadRepo,
		authStore: authStore,
	}

	m.load()

	return m
}

type Module struct {
	leadRepo  LeadRepository
	authStore auth.AuthorizationStore
}

func (m *Module) load() {
	m.authStore.RegisterPermission(auth.Permission{Id: PermissionLeadsView, Name: "Can view leads"})
}
