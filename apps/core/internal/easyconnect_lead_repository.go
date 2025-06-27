package internal

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/scan"

	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect"
	"github.com/piksar-eu/webapp/apps/core/pkg/shared"
)

func NewPgEasyConnectLeadRepository(db *sql.DB) easyconnect.LeadRepository {
	return &pgEasyConnectLeadRepository{
		db: bob.NewDB(db),
	}
}

type pgEasyConnectLeadRepository struct {
	db bob.DB
}

func (r *pgEasyConnectLeadRepository) Get(id string) (*easyconnect.LeadSnapshot, error) {

	ctx := context.Background()

	q := r.selectLead(sm.Where(psql.Quote("id").EQ(psql.S(id))))

	lead, err := bob.One(ctx, r.db, q, scan.StructMapper[easyconnect.LeadSnapshot]())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // or return a custom not found error
		}
		return nil, err
	}

	return &lead, nil
}

func (r *pgEasyConnectLeadRepository) GetByEmail(email string) (*easyconnect.LeadSnapshot, error) {
	ctx := context.Background()

	q := r.selectLead(sm.Where(psql.Quote("email").EQ(psql.S(email))))

	lead, err := bob.One(ctx, r.db, q, scan.StructMapper[easyconnect.LeadSnapshot]())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // or return a custom not found error
		}
		return nil, err
	}

	return &lead, nil
}

func (r *pgEasyConnectLeadRepository) List() ([]easyconnect.LeadSnapshot, error) {

	ctx := context.Background()

	q := r.selectLead(nil)

	leads, err := bob.All(ctx, r.db, q, scan.StructMapper[easyconnect.LeadSnapshot]())
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (r *pgEasyConnectLeadRepository) Save(l *easyconnect.LeadSnapshot) error {

	ctx := context.Background()

	q := psql.Insert(
		im.IntoAs("easyconnect__leads", "l", "id", "email", "source", "marketing_consent", "created_at"),
		im.Values(psql.Arg(l.Id, l.Email, l.Source, l.MarketingConsent, l.CreatedAt)),
		im.OnConflict("id").DoUpdate(
			im.SetCol("email").To(psql.Raw("EXCLUDED.email")),
			im.SetCol("marketing_consent").To(psql.Raw("EXCLUDED.marketing_consent")),
		),
	)

	_, err := bob.Exec(ctx, r.db, q)

	if err != nil {
		return err
	}

	return nil
}

func (r *pgEasyConnectLeadRepository) NewId() string {
	for {
		id, _ := shared.RandId("el", 8)
		exists, _ := r.Get(id)
		if exists == nil {
			return id
		}
	}
}

func (r *pgEasyConnectLeadRepository) selectLead(where bob.Mod[*dialect.SelectQuery]) bob.BaseQuery[*dialect.SelectQuery] {
	sq := []bob.Mod[*dialect.SelectQuery]{
		sm.Columns("id", "email", "source", "marketing_consent", "created_at"),
		sm.From("easyconnect__leads"),
	}

	if where != nil {
		sq = append(sq, where)
	}

	return psql.Select(sq...)
}
