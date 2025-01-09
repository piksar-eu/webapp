package infrastructure

import (
	"database/sql"
	"errors"
	"time"

	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect"
)

func NewPgEasyConnectLeadRepository(db *sql.DB) easyconnect.LeadRepository {
	return &pgEasyConnectLeadRepository{
		db: db,
	}
}

type pgEasyConnectLeadRepository struct {
	db *sql.DB
}

func (r *pgEasyConnectLeadRepository) Get(email string) (*easyconnect.Lead, error) {
	rows, err := r.db.Query("SELECT source, marketing_consent, created_at FROM easyconnect__leads WHERE email = $1 LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("lead not found")
	}

	var source string
	var marketingConsent bool
	var createdAt time.Time

	err = rows.Scan(&source, &createdAt)
	if err != nil {
		return nil, err
	}

	return &easyconnect.Lead{
		Email:            email,
		Source:           source,
		MarketingConsent: marketingConsent,
		CreatedAt:        createdAt,
	}, nil
}

func (r *pgEasyConnectLeadRepository) Save(lead *easyconnect.Lead) error {
	query := `
		INSERT INTO easyconnect__leads (email, source, marketing_consent, created_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (email) DO UPDATE
		SET source = EXCLUDED.source,
			marketing_consent = EXCLUDED.marketing_consent,
			created_at = EXCLUDED.created_at;
	`
	_, err := r.db.Exec(query, lead.Email, lead.Source, lead.MarketingConsent, lead.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
