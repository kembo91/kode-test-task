package database

import (
	"database/sql"
	"fmt"

	"github.com/kembo91/kode-test-task/server/utils"
	"golang.org/x/crypto/bcrypt"
)

//Credentials is a struct for sorage and extraction of user data
type Credentials struct {
	Username string `json:"Username" db:"username"`
	Password string `json:"Password" db:"passwordhash"`
}

var insertUserStmt = `
	INSERT INTO users(
		username,
		passwordhash
	) VALUES (
		$1,
		$2
	)
`

var checkUserStmt = `
	SELECT * FROM users WHERE username = $1 LIMIT 1
`

//InsertUser inserst a user to database
func (d *Database) InsertUser(c Credentials) error {
	var cx Credentials
	err := d.db.Get(&cx, checkUserStmt, c.Username)
	switch err {
	case sql.ErrNoRows:
		break
	case nil:
		return error(fmt.Errorf("User with username %s already exists", c.Username))
	default:
		return err
	}
	err = utils.IsValidPassword(c.Password)
	if err != nil {
		return err
	}
	err = utils.IsValidUsername(c.Username)
	if err != nil {
		return err
	}
	pHash, err := hashPassword(c.Password)
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
	switch err {
	case sql.ErrNoRows:
		return fmt.Errorf(`User %v not found`, c.Username)
	case nil:
		break
	default:
		return err
	}
	err = comparePasswords(c.Password, u.Password)
	if err != nil {
		return fmt.Errorf(`Wrong password`)
	}
	return nil
}

//HashPassword hashes user password to store it in a database
func hashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//ComparePasswords compares hashed password from a database with provided password
func comparePasswords(pwd string, hashedPwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		return err
	}
	return nil
}
