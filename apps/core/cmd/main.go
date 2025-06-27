package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/piksar-eu/webapp/apps/core/internal/di"
	"github.com/piksar-eu/webapp/apps/core/internal/migrations"
	"github.com/piksar-eu/webapp/apps/core/pkg/auth"
	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect"
	_ "github.com/piksar-eu/webapp/apps/core/pkg/envloader"
	"github.com/piksar-eu/webapp/apps/core/pkg/traefik_auth"
	"github.com/piksar-eu/webapp/apps/core/pkg/web"
)

func init() {
	migrations.Migrate()
}

func main() {
	port, _ := strconv.Atoi(os.Getenv("WEBSITE_PORT"))
	go serveFrontend("website", port)
	port, _ = strconv.Atoi(os.Getenv("DASHBOARD_PORT"))
	go serveFrontend("dashboard", port)

	serveApi()
}

func serveApi() {
	port, _ := strconv.Atoi(os.Getenv("API_PORT"))

	mux := http.NewServeMux()

	easyconnectModule := easyconnect.LoadModule(di.NewLeadRepository(), di.NewAuthorizationStore())
	authModule := auth.LoadModule(di.NewUserRepository(), di.NewRoleRepository(), di.NewAuthorizationStore())
	traefikAuth := traefik_auth.LoadModule(di.NewAuthorizationStore())

	easyconnectModule.ServeApi(mux)
	authModule.ServeApi(mux)
	traefikAuth.ServeApi(mux)

	var handler http.Handler = mux
	handler = web.CorsMiddleware(handler)
	handler = auth.AuthorizationMiddleware(di.NewAuthorizationStore())(handler)
	handler = web.SessionMiddleware(di.NewSessionStore())(handler)

	log.Printf("Serve api on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	if err != nil {
		panic(err)
	}
}

func serveFrontend(app string, port int) {

	mux := http.NewServeMux()

	web.ServeUi(mux, app, userJsonData)

	var handler http.Handler = mux
	handler = web.SessionMiddleware(di.NewSessionStore())(handler)

	log.Printf("Serve %s on port %d", app, port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	if err != nil {
		panic(err)
	}
}

func userJsonData(r *http.Request) string {
	as := di.NewAuthorizationStore()

	if u := web.GetSessionUser(r); u != nil {
		up := as.UserPermissions(u.Id)
		for i, v := range up {
			up[i] = fmt.Sprintf(`"%s"`, v)
		}

		return fmt.Sprintf(`{
					email: "%s",
					permissions: [%s]
				}`, u.Email, strings.Join(up, ","))
	}

	return ""
}
