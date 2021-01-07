package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kembo91/kode-test-task/server/database"
)

//InsertAnagram handles anagram insertion request
func InsertAnagram(db *database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var q database.Query
		err := json.NewDecoder(r.Body).Decode(&q)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = db.InsertAnagram(strings.ToLower(q.QueryString))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}

//RetrieveAnagram handles anagram query retrieval request
func RetrieveAnagram(db *database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var q database.Query
		err := json.NewDecoder(r.Body).Decode(&q)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		anagrams, err := db.RetrieveQueryAnagram(strings.ToLower(q.QueryString))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(anagrams)
	}
}

//RetrieveAll handles all anagrams retrieval request
func RetrieveAll(db *database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var q database.Query
		err := json.NewDecoder(r.Body).Decode(&q)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		anagrams, err := db.RetrieveAllAnagrams()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(anagrams)
	}
}
