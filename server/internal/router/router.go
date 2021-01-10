package router

import (
	"net/http"

	"github.com/kembo91/kode-test-task/server/internal/database"
	"github.com/kembo91/kode-test-task/server/internal/handlers"

	"github.com/gorilla/mux"
)

//CreateRouter creates a mux router and sets up routes
func CreateRouter(db *database.Database) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/signin", handlers.SigninHandler(db)).Methods("POST")
	api.HandleFunc("/signup", handlers.SignupHandler(db)).Methods("POST")
	api.HandleFunc("/signout", handlers.SignoutHandler).Methods("GET")

	anagrapi := api.PathPrefix("/anagram").Subrouter()
	anagrapi.Use(handlers.AuthenticationMiddleware)

	anagrapi.HandleFunc("/insert", handlers.InsertAnagram(db)).Methods("POST")
	anagrapi.HandleFunc("/retrieve", handlers.RetrieveAnagram(db)).Methods("POST")
	anagrapi.HandleFunc("/retrieve", handlers.RetrieveAll(db)).Methods("GET")

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static")))
	r.PathPrefix("/static/").Handler(staticHandler)

	buildHandler := http.FileServer(http.Dir("./build"))
	r.PathPrefix("/").Handler(buildHandler)
	return r
}
