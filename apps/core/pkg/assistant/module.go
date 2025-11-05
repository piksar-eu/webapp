package assistant

import (
	"github.com/piksar-eu/webapp/apps/core/pkg/assistant/public"
	"github.com/piksar-eu/webapp/apps/core/pkg/command"
)

func LoadModule(
	commandBus *command.CommandBus,
) *Module {
	m := &Module{
		commandBus: commandBus,
	}

	m.load()

	return m
}

type Module struct {
	commandBus *command.CommandBus
}

func (m *Module) load() {
	m.commandBus.RegisterHandler(public.CallAssistantCommand{}, &CallAssistantHandler{})
}
