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
				id VARCHAR PRIMARY KEY,
				email VARCHAR(255) NOT NULL,
				name VARCHAR NOT NULL DEFAULT '',
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				auth_methods JSONB
			);`); err != nil {
				return err
			}

			return nil
		},
	}
}
