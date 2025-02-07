package migrations

import (
	"github.com/piksar-eu/webapp/apps/core/pkg/di"
)

func m202501270750_users() Migration {
	return Migration{
		Id: "202501270750_users",
		Up: func() error {
			db := di.NewDb()

			if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS auth__users (
				email VARCHAR(255) PRIMARY KEY,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				auth_methods JSONB
			);`); err != nil {
				return err
			}

			return nil
		},
	}
}
