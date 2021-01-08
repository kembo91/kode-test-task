package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kembo91/kode-test-task/server/internal/utils"

	"github.com/kembo91/kode-test-task/server/internal/database"
)

//InsertAnagram handles anagram insertion request
func InsertAnagram(db *database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var q database.Query
		err := json.NewDecoder(r.Body).Decode(&q)
		if err != nil {
			utils.JSONError(w, err, http.StatusBadRequest)
			return
		}
		if q.QueryString == "" {
			err = fmt.Errorf(`expected a non-empty field "Query"`)
			utils.JSONError(w, err, http.StatusBadRequest)
			return
		}
		err = db.InsertAnagram(strings.ToLower(q.QueryString))
		if err != nil {
			utils.JSONError(w, err, http.StatusInternalServerError)
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
			utils.JSONError(w, err, http.StatusBadRequest)
			return
		}
		if q.QueryString == "" {
			err = fmt.Errorf(`expected a non-empty field "Query"`)
			utils.JSONError(w, err, http.StatusBadRequest)
			return
		}
		anagrams, err := db.RetrieveQueryAnagram(strings.ToLower(q.QueryString))
		if err != nil {
			utils.JSONError(w, err, http.StatusInternalServerError)
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
			utils.JSONError(w, err, http.StatusBadRequest)
			return
		}
		anagrams, err := db.RetrieveAllAnagrams()
		if err != nil {
			utils.JSONError(w, err, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(anagrams)
	}
}
