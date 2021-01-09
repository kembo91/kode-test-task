package database_test

import (
	"testing"

	"github.com/kembo91/kode-test-task/server/internal/database"
)

func TestInsertUserCorrect(t *testing.T) {
	usr := database.Credentials{Username: "hello", Password: "worldherebro"}
	err := db.InsertUser(usr)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertExistingUser(t *testing.T) {
	usr := database.Credentials{Username: "hellog", Password: "worldherebro"}
	err := db.InsertUser(usr)
	if err != nil {
		t.Error(err)
	}
	err = db.InsertUser(usr)
	if err == nil {
		t.Error(err)
	}
}

func TestInsertUserWrongUsername(t *testing.T) {
	usr := database.Credentials{Username: "he", Password: "worldherebro"}
	err := db.InsertUser(usr)
	if err == nil {
		t.Error(err)
	}
}

func TestInsertUserWrongPass(t *testing.T) {
	usr := database.Credentials{Username: "helloworld", Password: "wor"}
	err := db.InsertUser(usr)
	if err == nil {
		t.Error(err)
	}
}

func TestCheckUserWrongPass(t *testing.T) {
	usr := database.Credentials{Username: "hellof", Password: "worldherebro"}
	err := db.InsertUser(usr)
	if err != nil {
		t.Error(err)
	}
	usr.Password = "wrongpass"
	err = db.CheckUser(usr)
	if err == nil {
		t.Errorf("User password check failed")
	}
}

func TestCheckUserWrongUsername(t *testing.T) {
	usr := database.Credentials{Username: "hellod", Password: "worldherebro"}
	err := db.InsertUser(usr)
	if err != nil {
		t.Error(err)
	}
	usr.Username = "wrongusername"
	err = db.CheckUser(usr)
	if err == nil {
		t.Errorf("Username check failed")
	}
}
