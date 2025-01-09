package easyconnect

import (
	"time"

	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

type LeadRepository interface {
	Get(email string) (*Lead, error)
	Save(*Lead) error
}

type Lead struct {
	Email            string
	MarketingConsent bool
	Source           string
	CreatedAt        time.Time
}

func createLead(email string, source string) (*Lead, error) {

	email, err := shared.SanitizeEmail(email)

	if err != nil {
		return nil, err
	}

	return &Lead{
		Email:            email,
		MarketingConsent: false,
		Source:           source,
		CreatedAt:        time.Now(),
	}, nil
}

func (l *Lead) changeMarketingConsent(consent bool) {
	l.MarketingConsent = consent
}
