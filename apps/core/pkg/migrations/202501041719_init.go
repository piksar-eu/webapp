package migrations

import (
	"log"

	"github.com/piksar-eu/webapp/apps/core/pkg/di"
)

func m202501041719_init() Migration {
	return Migration{
		Id: "202501041719_init",
		Up: func() error {

			db := di.NewDb()

			_, err := db.Exec(`CREATE TABLE IF NOT EXISTS core__migration_log (
				id SERIAL PRIMARY KEY,
				date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				name VARCHAR(255),
				err TEXT
			);`)
			if err != nil {
				panic(err)
			}

			log.Println("Created migration log table")

			return nil
		},
	}
}
