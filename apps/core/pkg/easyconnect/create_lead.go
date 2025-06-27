package easyconnect

import (
	"context"
	"fmt"

	"github.com/piksar-eu/webapp/apps/core/pkg/command"
	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect/public"
	"github.com/piksar-eu/webapp/apps/core/pkg/events"
	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

type CreateLeadHandler struct {
	leadRepo    LeadRepository
	evPublisher events.EventPublisher
}

func (h *CreateLeadHandler) Handle(ctx context.Context, cm command.Command) (any, error) {
	c, ok := cm.(public.CreateLeadCommand)
	if !ok {
		return nil, fmt.Errorf("invalid command type")
	}

	email, err := shared.SanitizeEmail(c.Email)

	if err != nil {
		return nil, err
	}

	var lead *Lead

	leadSnap, err := h.leadRepo.GetByEmail(email)
	if leadSnap != nil {
		lead, err = leadFromSnapshot(leadSnap)
	}

	if err != nil {
		return nil, err
	}

	if lead != nil && lead.MarketingConsent != c.MarketingConsent {
		return nil, nil
	}

	if lead == nil {
		id := h.leadRepo.NewId()

		lead, err = createLead(id, email, c.Source)
	}

	if err != nil {
		return nil, err
	}

	lead.changeMarketingConsent(c.MarketingConsent)

	if err := h.leadRepo.Save(lead.Snapshot()); err != nil {
		return nil, fmt.Errorf("failed to save Lead: %w", err)
	}

	h.evPublisher.Publish(lead.PullEvents()...)

	return &public.CreateLeadRes{Id: lead.Id}, nil
}
