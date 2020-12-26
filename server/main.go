package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kembo91/kode-test-task/server/handlers/userauth"

	"github.com/kembo91/kode-test-task/server/database"

	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.CreateDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(db)
	r := mux.NewRouter()
	buildHandler := http.FileServer(http.Dir("./build"))
	r.PathPrefix("/").Handler(buildHandler)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static")))
	r.PathPrefix("/static/").Handler(staticHandler)

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/login", userauth.SigninHandler(db)).Methods("POST")
	api.HandleFunc("/signup", userauth.SignupHandler(db)).Methods("POST")

	anagramm := api.PathPrefix("/anagramm").Subrouter()

	anagramm.Use(userauth.AuthenticationMiddleware)

	srv := http.Server{
		Handler:      h.LoggingHandler(os.Stdout, r),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
