package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
	go serveWebsite()
	serveApi()
}

func serveApi() {
	port, _ := strconv.Atoi(os.Getenv("API_PORT"))

	mux := http.NewServeMux()

	easyconnect.ServeApi(mux, di.NewLeadRepository())

	var handler http.Handler = mux
	handler = web.CorsMiddleware(handler)

	log.Printf("Serve api on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	if err != nil {
		panic(err)
	}
}

func serveWebsite() {
	port, _ := strconv.Atoi(os.Getenv("WEBSITE_PORT"))

	mux := http.NewServeMux()

	web.ServeUi(mux)

	var handler http.Handler = mux

	log.Printf("Serve website on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	if err != nil {
		panic(err)
	}
}
