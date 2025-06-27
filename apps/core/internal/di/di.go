package di

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"github.com/piksar-eu/webapp/apps/core/internal"
	"github.com/piksar-eu/webapp/apps/core/pkg/auth"
	"github.com/piksar-eu/webapp/apps/core/pkg/command"
	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect"
	"github.com/piksar-eu/webapp/apps/core/pkg/events"
	"github.com/piksar-eu/webapp/apps/core/pkg/web"
)

var services = struct {
	DB                 *sql.DB
	LeadRepository     easyconnect.LeadRepository
	SessionStore       web.SessionStore
	UserRepository     auth.UserRepository
	RoleRepository     auth.RoleRepository
	AuthorizationStore auth.AuthorizationStore
	CommandBus         *command.CommandBus
	EventPublisher     events.EventPublisher
}{}

var dbOnce sync.Once

func NewDb() *sql.DB {
	dbOnce.Do(func() {
		connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DBNAME"))
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		services.DB = db
	})

	return services.DB
}

var commandbusOnce sync.Once

func NewCommandBus() *command.CommandBus {
	commandbusOnce.Do(func() {
		services.CommandBus = command.NewCommandBus()
	})

	return services.CommandBus
}

var leadRepositoryOnce sync.Once

func NewLeadRepository() easyconnect.LeadRepository {
	leadRepositoryOnce.Do(func() {
		services.LeadRepository = internal.NewPgEasyConnectLeadRepository(NewDb())
	})

	return services.LeadRepository
}

var userRepositoryOnce sync.Once

func NewUserRepository() auth.UserRepository {
	userRepositoryOnce.Do(func() {
		services.UserRepository = internal.NewPgAuthUserRepository(NewDb())
	})

	return services.UserRepository
}

var roleRepositoryOnce sync.Once

func NewRoleRepository() auth.RoleRepository {
	roleRepositoryOnce.Do(func() {
		services.RoleRepository = internal.NewPgAuthRoleRepository(NewDb())
	})

	return services.RoleRepository
}

var sessionStoreOnce sync.Once

func NewSessionStore() web.SessionStore {
	sessionStoreOnce.Do(func() {
		pgSessionStore := internal.NewPgSessionStore(NewDb())
		services.SessionStore = internal.NewCachedSessionStore(pgSessionStore)
	})

	return services.SessionStore
}

var authorizationStoreOnce sync.Once

func NewAuthorizationStore() auth.AuthorizationStore {
	authorizationStoreOnce.Do(func() {
		services.AuthorizationStore = internal.NewAuthorizationStore(NewUserRepository(), NewRoleRepository())
	})

	return services.AuthorizationStore
}

var eventPublisherOnce sync.Once

func NewEventPublisher() events.EventPublisher {
	eventPublisherOnce.Do(func() {
		services.EventPublisher = internal.NewSyncEventPublisher()
	})

	return services.EventPublisher
}
