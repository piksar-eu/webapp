package easyconnect

import (
	"github.com/piksar-eu/webapp/apps/core/pkg/auth"
	"github.com/piksar-eu/webapp/apps/core/pkg/command"
	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect/public"
	"github.com/piksar-eu/webapp/apps/core/pkg/events"
)

const PermissionLeadsView = "leads_view"

func LoadModule(
	leadRepo LeadRepository,
	authStore auth.AuthorizationStore,
	commandBus *command.CommandBus,
	evPublisher events.EventPublisher,
) *Module {
	m := &Module{
		leadRepo:    leadRepo,
		authStore:   authStore,
		commandBus:  commandBus,
		evPublisher: evPublisher,
	}

	m.load()

	return m
}

type Module struct {
	leadRepo    LeadRepository
	authStore   auth.AuthorizationStore
	commandBus  *command.CommandBus
	evPublisher events.EventPublisher
}

func (m *Module) load() {
	m.authStore.RegisterPermission(auth.Permission{Id: PermissionLeadsView, Name: "Can view leads"})

	m.commandBus.RegisterHandler(public.CreateLeadCommand{}, &CreateLeadHandler{
		leadRepo:    m.leadRepo,
		evPublisher: m.evPublisher,
	})
}
