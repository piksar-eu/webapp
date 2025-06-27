package public

import (
	"github.com/piksar-eu/webapp/apps/core/pkg/command"
)

func NewCreateLeadCommand(email, source string, marketingConsent bool) CreateLeadCommand {
	cmd := CreateLeadCommand{
		Email:            email,
		Source:           source,
		MarketingConsent: marketingConsent,
	}

	cmd.BaseCommand = command.NewBaseCommand(cmd, nil, nil)

	return cmd
}

type CreateLeadCommand struct {
	command.BaseCommand
	Email            string `json:"email"`
	Source           string `json:"source"`
	MarketingConsent bool   `json:"marketingConsent"`
}

type CreateLeadRes struct {
	Id string `json:"id"`
}
