package migrations

import (
	"github.com/piksar-eu/webapp/apps/core/pkg/di"
)

func m202501162300_sessions() Migration {
	return Migration{
		Id: "202501162300_sessions",
		Up: func() error {
			db := di.NewDb()

			if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS core__sessions (
				id UUID PRIMARY KEY,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				expires_at TIMESTAMP NOT NULL,
				data JSONB
			);`); err != nil {
				return err
			}

			return nil
		},
	}
}
