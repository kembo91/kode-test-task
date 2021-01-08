package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kembo91/kode-test-task/server/internal/utils"

	"github.com/kembo91/kode-test-task/server/internal/router"

	"github.com/kembo91/kode-test-task/server/internal/database"

	h "github.com/gorilla/handlers"
)

func main() {
	cfgPath := "../../config/dbconfig.yaml"
	cfg := utils.GetDBConfig(cfgPath)
	db, err := database.CreateDB("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}
	r := router.CreateRouter(db)
	srv := http.Server{
		Handler:      h.LoggingHandler(os.Stdout, r),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
