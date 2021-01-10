package main

import (
	"fmt"
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
	cfg := utils.GetDBConfig()
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
	fmt.Println("Now listening on port 8080")
	srv.ListenAndServe()
}
