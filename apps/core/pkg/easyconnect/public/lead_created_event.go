package public

import (
	"github.com/piksar-eu/webapp/apps/core/pkg/events"
)

const LeadCreatedEvent = "easyconnect.leadCreated"

type LeadCreatedPayload struct {
	LeadId string `json:"leadId"`
}

func NewLeadCreatedEvent(leadId string) events.Event {
	return events.NewEvent(LeadCreatedEvent, LeadCreatedPayload{
		LeadId: leadId,
	})
}
