package database

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kembo91/kode-test-task/server/utils"
)

//Query struct for parsing incomig queries
type Query struct {
	QueryString string `json:"Query"`
}

var getAnagramStmt = `
	SELECT anagrams_id FROM anagrams WHERE anagram = $1
`

var insertAnagramStmt = `
	INSERT INTO anagrams (sorted_id, anagram) VALUES ($1, $2)
`

var insertSortedAnagramStmt = `
	INSERT INTO sortedanagrams (sorted) VALUES ($1) RETURNING sorted_id
`

var getSortedAnagramStmt = `
	SELECT sorted_id FROM sortedanagrams WHERE Sorted = $1
`

var getQueryAnagramStmt = `
	SELECT DISTINCT anagrams.anagram
	FROM sortedanagrams
	LEFT JOIN anagrams
	ON (
		SELECT sorted_id FROM sortedanagrams WHERE sorted=$1
	) = anagrams.sorted_id
`

var getAllAnagramsStmt = `
	SELECT DISTINCT anagram FROM anagrams
`

//InsertAnagram checks anagram existence in the database and inserts in case it doesn't exist
func (d *Database) InsertAnagram(a string) error {
	err := utils.IsValidAnagram(a)
	if err != nil {
		return err
	}
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
	anagrams := make([]string, 0)
	err := utils.IsValidAnagram(a)
	if err != nil {
		return anagrams, err
	}
	sorted := d.sortAnagram(a)
	err = d.db.Select(&anagrams, getQueryAnagramStmt, sorted)
	if err != nil {
		return anagrams, err
	}
	return anagrams, nil
}

//RetrieveAllAnagrams retrieves all anagrams from database
func (d *Database) RetrieveAllAnagrams() ([]string, error) {
	var anagrams []string
	err := d.db.Select(&anagrams, getAllAnagramsStmt)
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
