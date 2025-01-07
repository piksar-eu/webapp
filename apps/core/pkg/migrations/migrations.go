package migrations

import (
	"log"

	"github.com/piksar-eu/webapp/apps/core/pkg/di"
)

type Migration struct {
	Id string
	Up func() error
}

func Migrate() {
	migrationLog := migrationLog()

	for _, migration := range migrations() {
		if isMigrated(migration.Id, migrationLog) {
			continue
		}

		err := migration.Up()

		logMigration(migration.Id)

		if err != nil {
			logMigrationError(migration.Id, err)
		}
	}
}

func migrations() []Migration {
	return []Migration{
		m202501041719_init(),
	}
}

func migrationLog() []string {
	db := di.NewDb()

	rows, err := db.Query("SELECT name FROM core__migration_log WHERE err IS NULL")
	if err != nil {
		log.Println(" no migration log yet...")
		return []string{}
	}
	defer rows.Close()

	arr := []string{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}

		arr = append(arr, name)
	}

	return arr
}

func isMigrated(migrationId string, log []string) bool {
	for _, name := range log {
		if name == migrationId {
			return true
		}
	}
	return false
}

func logMigration(name string) {
	db := di.NewDb()
	_, err := db.Exec("INSERT INTO core__migration_log (name) VALUES ($1);", name)
	if err != nil {
		panic(err)
	}

	log.Printf("Migrated %s", name)
}

func logMigrationError(name string, migrationErr error) {
	db := di.NewDb()
	_, err := db.Exec("UPDATE core__migration_log SET err = $1 WHERE name = $2", migrationErr.Error(), name)
	if err != nil {
		panic(err)
	}
}
