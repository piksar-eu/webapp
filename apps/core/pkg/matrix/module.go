package matrix

import (
	"database/sql"

	"github.com/piksar-eu/webapp/apps/core/pkg/command"
	"github.com/piksar-eu/webapp/apps/core/pkg/events"
)

type Module struct {
	db          *sql.DB
	matrix      *Matrix
	evPublisher events.EventPublisher
	commandBus  *command.CommandBus
}

func (m *Module) load() {
	m.matrix, _ = NewMatrix(m.db, m.commandBus)

	// go m.matrix.syncMatrix()

	// m.evPublisher.Register(easyconnect_public.LeadCreatedEvent, m.SendNotificationOnLeadCreatedListener)
}

func LoadModule(
	db *sql.DB,
	evPublisher events.EventPublisher,
	commandBus *command.CommandBus,
) *Module {
	m := &Module{
		db:          db,
		evPublisher: evPublisher,
		commandBus:  commandBus,
	}

	m.load()

	return m
}
