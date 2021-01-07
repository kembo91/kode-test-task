package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	//postgresql database driver
	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "postgres"
)

//Database is a basic database struct to encapsulate functions
type Database struct {
	db *sqlx.DB
}

var createUsersTableStmt = `
	CREATE TABLE IF NOT EXISTS users(
		username			TEXT,
		passwordhash		TEXT
	)
`
var createSortedAnagramTableStmt = `
	CREATE TABLE IF NOT EXISTS sortedanagrams(
		sorted_id				SERIAL,
		sorted	TEXT,
		PRIMARY KEY (sorted_id)
	)
`

var createAnagramTableStmt = `
	CREATE TABLE IF NOT EXISTS anagrams(
		anagrams_id				SERIAL,
		sorted_id		int NOT NULL,
		anagram			TEXT,
		PRIMARY KEY (anagrams_id),
		FOREIGN KEY (sorted_id) REFERENCES sortedanagrams(sorted_id) ON DELETE CASCADE
	)
`

//CreateDB creates a database and sets up tables if they don't exist
func CreateDB() (*Database, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createUsersTableStmt)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createSortedAnagramTableStmt)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createAnagramTableStmt)
	if err != nil {
		return nil, err
	}
	var d Database
	d.db = db
	return &d, nil
}
