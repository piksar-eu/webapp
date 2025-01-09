package di

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect"
	"github.com/piksar-eu/webapp/apps/core/pkg/infrastructure"
)

var services = struct {
	DB             *sql.DB
	LeadRepository easyconnect.LeadRepository
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

func NewLeadRepository() easyconnect.LeadRepository {
	if services.LeadRepository == nil {
		services.LeadRepository = infrastructure.NewPgEasyConnectLeadRepository(NewDb())
	}

	return services.LeadRepository
}
