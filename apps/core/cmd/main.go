package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/piksar-eu/webapp/apps/core/pkg/auth"
	"github.com/piksar-eu/webapp/apps/core/pkg/di"
	"github.com/piksar-eu/webapp/apps/core/pkg/easyconnect"
	_ "github.com/piksar-eu/webapp/apps/core/pkg/envloader"
	"github.com/piksar-eu/webapp/apps/core/pkg/migrations"
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

	easyconnect.ServeApi(mux, di.NewLeadRepository())
	auth.ServeApi(mux, di.NewUserRepository())

	var handler http.Handler = mux
	handler = web.CorsMiddleware(handler)
	handler = web.SessionMiddleware(di.NewSessionStore())(handler)

	log.Printf("Serve api on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	if err != nil {
		panic(err)
	}
}

func serveFrontend(app string, port int) {

	mux := http.NewServeMux()

	web.ServeUi(mux, app)

	var handler http.Handler = mux
	handler = web.SessionMiddleware(di.NewSessionStore())(handler)

	log.Printf("Serve %s on port %d", app, port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	if err != nil {
		panic(err)
	}
}
