package easyconnect

import (
	"time"

	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect/public"
	"github.com/piksar-eu/webapp/apps/core/pkg/events"
	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

type LeadRepository interface {
	Get(id string) (*LeadSnapshot, error)
	GetByEmail(email string) (*LeadSnapshot, error)
	List() ([]LeadSnapshot, error)
	Save(*LeadSnapshot) error
	NewId() string
}

type LeadSnapshot struct {
	Id               string    `json:"id"`
	Email            string    `json:"email"`
	Source           string    `json:"source"`
	MarketingConsent bool      `json:"marketingConsent"`
	CreatedAt        time.Time `json:"createdAt"`
}

type Lead struct {
	Id               string
	Email            string
	Source           string
	MarketingConsent bool
	CreatedAt        time.Time

	events.DomainEvents
}

func createLead(id, email, source string) (*Lead, error) {
	email, err := shared.SanitizeEmail(email)

	if err != nil {
		return nil, err
	}

	lead := &Lead{
		Id:               id,
		Email:            email,
		MarketingConsent: false,
		Source:           source,
		CreatedAt:        time.Now(),
	}

	lead.Record(public.NewLeadCreatedEvent(id))

	return lead, nil
}

func leadFromSnapshot(s *LeadSnapshot) (*Lead, error) {
	return &Lead{
		Id:               s.Id,
		Email:            s.Email,
		Source:           s.Source,
		MarketingConsent: s.MarketingConsent,
		CreatedAt:        s.CreatedAt,
	}, nil
}

func (i *Lead) Snapshot() *LeadSnapshot {
	return &LeadSnapshot{
		Id:               i.Id,
		Email:            i.Email,
		Source:           i.Source,
		MarketingConsent: i.MarketingConsent,
		CreatedAt:        i.CreatedAt,
	}
}

func (l *Lead) changeMarketingConsent(consent bool) {
	l.MarketingConsent = consent
}
