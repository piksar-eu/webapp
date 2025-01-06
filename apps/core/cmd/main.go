package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/piksar-eu/webapp/apps/core/pkg/envloader"
	"github.com/piksar-eu/webapp/apps/core/pkg/migrations"
)

func init() {
	migrations.Migrate()
}

func main() {
	apiPort, _ := strconv.Atoi(os.Getenv("API_PORT"))

	mux := http.NewServeMux()

	log.Printf("Serve app on port %d", apiPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", apiPort), mux)

	if err != nil {
		panic(err)
	}
}
