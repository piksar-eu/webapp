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
	apiPort, _ := strconv.Atoi(os.Getenv("API_PORT"))

	mux := http.NewServeMux()

	easyconnect.ServeApi(mux, di.NewLeadRepository())

	var handler http.Handler = mux
	handler = web.CorsMiddleware(handler)

	log.Printf("Serve app on port %d", apiPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", apiPort), handler)

	if err != nil {
		panic(err)
	}
}
