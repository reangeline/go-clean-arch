package http

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/reangeline/go-clean-arch/config"
	"github.com/reangeline/go-clean-arch/internal/di"
	"github.com/reangeline/go-clean-arch/internal/infra/http/route"
)

func ServerHttp(db *sql.DB, config *config.Conf) {
	router := chi.NewRouter()

	gor, err := di.InitializeOrder(db)
	if err != nil {
		log.Fatalf("failed to initialize order controller: %v", err)
	}

	route.InitializeOrderRoutes(gor, router)

	log.Printf("connect to http://localhost:%s/ for Rest Api", config.WebServerPort)
	err = http.ListenAndServe(":"+config.WebServerPort, router)
	if err != nil {
		panic(err)
	}

}
