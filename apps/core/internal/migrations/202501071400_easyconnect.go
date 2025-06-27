package migrations

import (
	"github.com/piksar-eu/webapp/apps/core/internal/di"
)

func m202501071400_easyconnect() Migration {
	return Migration{
		Id: "202501071400_easyconnect",
		Up: func() error {
			db := di.NewDb()

			if _, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS easyconnect__leads (
				id TEXT PRIMARY KEY,
				email TEXT,
				source TEXT,
				marketing_consent BOOLEAN,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			);`); err != nil {
				return err
			}

			return nil
		},
	}
}
