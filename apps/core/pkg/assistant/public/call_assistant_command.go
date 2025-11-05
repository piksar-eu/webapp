package public

import (
	"github.com/piksar-eu/webapp/apps/core/pkg/command"
)

func NewCallAssistantCommand(content string) CallAssistantCommand {
	cmd := CallAssistantCommand{
		Content: content,
	}

	cmd.BaseCommand = command.NewBaseCommand(cmd, nil, nil)

	return cmd
}

type CallAssistantCommand struct {
	command.BaseCommand
	Content string `json:"content"`
}

type CallAssistantRes struct {
	Content string `json:"message,omitempty"`
}
