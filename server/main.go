package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kembo91/kode-test-task/server/database"
	"github.com/kembo91/kode-test-task/server/handlers"

	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.CreateDB()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/signin", handlers.SigninHandler(db)).Methods("POST")
	api.HandleFunc("/signup", handlers.SignupHandler(db)).Methods("POST")

	anagrapi := api.PathPrefix("/anagram").Subrouter()
	anagrapi.Use(handlers.AuthenticationMiddleware)

	anagrapi.HandleFunc("/insert", handlers.InsertAnagram(db)).Methods("POST")
	anagrapi.HandleFunc("/retrieve", handlers.RetrieveAnagram(db)).Methods("POST")
	anagrapi.HandleFunc("/retrieve", handlers.RetrieveAll(db)).Methods("GET")

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static")))
	r.PathPrefix("/static/").Handler(staticHandler)

	buildHandler := http.FileServer(http.Dir("./build"))
	r.PathPrefix("/").Handler(buildHandler)

	srv := http.Server{
		Handler:      h.LoggingHandler(os.Stdout, r),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
