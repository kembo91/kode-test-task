package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	//postgresql database driver
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "postgres"
)

//Credentials is a struct for sorage and extraction of user data
type Credentials struct {
	Username string `json:"Username" db:"Username"`
	Password string `json:"Password" db:"Password"`
}

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
	SELECT * FROM Users WHERE Username = ? LIMIT 1
`

//HashPassword hashes user password to store it in a database
func HashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//ComparePasswords compares hashed password from a database with provided password
func ComparePasswords(hashedPwd string, pwd string) error {
	bytehashed := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(bytehashed, []byte(pwd))
	if err != nil {
		return err
	}
	return nil
}

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

//InsertUser inserst a user to database
func (d *Database) InsertUser(c Credentials) error {
	_, err := d.db.Queryx(checkUserStmt, c.Username)
	switch err {
	case sql.ErrNoRows:
		break
	case nil:
		return error(fmt.Errorf("User with username %s already exists", c.Username))
	default:
		return err
	}
	pHash, err := HashPassword(c.Password)
	if err != nil {
		return err
	}
	tx := d.db.MustBegin()
	tx.MustExec(insertUserStmt, c.Username, pHash)
	tx.Commit()
	return nil
}

//CheckUser checks user existence and verifies password information
func (d *Database) CheckUser(c Credentials) error {
	var u Credentials
	err := d.db.Get(&u, checkUserStmt, c.Username)
	if err != nil {
		return err
	}
	hPwd, err := HashPassword(c.Password)
	if err != nil {
		return err
	}
	err = ComparePasswords(hPwd, u.Password)
	if err != nil {
		return err
	}
	return nil
}
