package main

import (
	"database/sql"
	"fmt"

	"github.com/reangeline/go-clean-arch/config"
	"github.com/reangeline/go-clean-arch/internal/infra/http"

	_ "github.com/lib/pq"
)

func main() {
	configs, err := config.LoadConfig(".")
	if err != nil {
		panic(configs)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	http.ServerHttp(db, configs)

}
