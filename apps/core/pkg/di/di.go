package di

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var services = struct {
	DB *sql.DB
}{}

func NewDb() *sql.DB {
	if services.DB == nil {
		connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DBNAME"))
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		services.DB = db
	}

	return services.DB
}
