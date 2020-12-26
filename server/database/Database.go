package database

import (
	"database/sql"
	"fmt"

	"github.com/kembo91/kode-test-task/server/handlers/userauth"

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
	CREATE TABLE IF NOT EXISTS Users(
		Username			TEXT,
		PasswordHash		TEXT
	)
`

var insertUserStmt = `
	INSERT INTO Users(
		Username,
		PasswordHash
	) VALUES (
		$1,
		$2
	)
`

var checkUserStmt = `
	SELECT Username FROM Users WHERE Username = ? LIMIT 1
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
	var d Database
	d.db = db
	return &d, nil
}

func (d *Database) InsertUser(c userauth.Credentials) error {
	_, err := d.db.Queryx(checkUserStmt, c.Username)
	switch err {
	case sql.ErrNoRows:
		return err
	case nil:
		break
	default:
		return err
	}

	return nil
}
