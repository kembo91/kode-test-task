package database

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

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

//Query struct for parsing incomig queries
type Query struct {
	QueryString string `json:"Query"`
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
var createSortedAnagramTableStmt = `
	CREATE TABLE IF NOT EXISTS SortedAnagrams(
		sorted_id				SERIAL,
		Sorted	TEXT,
		PRIMARY KEY (id)
	)
`

var createAnagramTableStmt = `
	CREATE TABLE IF NOT EXISTS Anagrams(
		anagrams_id				SERIAL,
		sorted_id		int NOT NULL,
		Anagram			TEXT,
		PRIMARY KEY (anagrams_id),
		FOREIGN KEY (sorted_id) REFERENCES SortedAnagrams(sorted_id) ON DELETE CASCADE
	)
`

var getAnagramStmt = `
	SELECT id FROM Anagrams WHERE Anagram = $1
`

var insertAnagramStmt = `
	INSERT INTO Anagrams (sorted_id, Anagram) VALUES ($1, $2)
`

var insertSortedAnagramStmt = `
	INSERT INTO SortedAnagrams (Sorted) VALUES ($1) RETURNING id
`

var getSortedAnagramStmt = `
	SELECT id FROM SortedAnagrams WHERE Sorted = $1
`

var getQueryAnagramStmt = `
	SELECT DISTINCT Anagrams.Anagram
	FROM SortedAnagrams
	LEFT JOIN Anagrams
	ON (
		SELECT sorted_id FROM SortedAnagrams WHERE Sorted=$1
	) = Anagrams.sorted_id
`

var getAllAnagramsStmt = `
	SELECT DISTINCT Anagram FROM Anagrams
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
	_, err = db.Exec(createAnagramTableStmt)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createSortedAnagramTableStmt)
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

//InsertAnagram checks anagram existence in the database and inserts in case it doesn't exist
func (d *Database) InsertAnagram(a string) error {
	sorted := d.sortAnagram(a)
	var sortedID int
	rows, err := d.db.Query(getSortedAnagramStmt, sorted)
	if err != nil {
		return err
	}
	if rows.Next() {
		rows.Scan(&sortedID)
	} else {
		sortedID, err = d.insertSortedAnagram(sorted)
		if err != nil {
			return err
		}
	}

	rows, err = d.db.Query(getAnagramStmt, a)
	if err != nil {
		return err
	}
	if rows.Next() {
		return fmt.Errorf(`Anagram %v already exists in the database`, a)
	}
	_, err = d.db.Exec(insertAnagramStmt, sortedID, a)
	if err != nil {
		return err
	}
	return nil
}

//RetrieveQueryAnagram retrieves all anagrams for an incoming string from database
func (d *Database) RetrieveQueryAnagram(a string) ([]string, error) {
	sorted := d.sortAnagram(a)
	anagrams := make([]string, 0)
	err := d.db.Get(&anagrams, getQueryAnagramStmt, sorted)
	if err != nil {
		return anagrams, err
	}
	return anagrams, nil
}

//RetrieveAllAnagrams retrieves all anagrams from database
func (d *Database) RetrieveAllAnagrams() ([]string, error) {
	anagrams := make([]string, 0)
	err := d.db.Get(&anagrams, getAllAnagramsStmt)
	if err != nil {
		return anagrams, err
	}
	return anagrams, nil
}

func (d *Database) sortAnagram(a string) string {
	t := strings.Split(a, "")
	sort.Strings(t)
	return strings.Join(t, "")
}

func (d *Database) insertSortedAnagram(a string) (int, error) {
	rows, err := d.db.Query(insertSortedAnagramStmt, a)
	if err != nil {
		return 0, err
	}
	var id int
	if rows.Next() {
		rows.Scan(&id)
	}
	return id, nil
}
