package database_test

import (
	"fmt"
	"testing"

	"github.com/kembo91/kode-test-task/server/internal/utils"

	"github.com/DATA-DOG/go-txdb"
	"github.com/kembo91/kode-test-task/server/internal/database"
	_ "github.com/lib/pq"
)

var cfg = utils.GetDBConfig("../../../config/dbtestconfig.yaml")

var insertTestStmt = `
	SELECT username from Users WHERE username = $1
`

func init() {
	txdb.Register("psql_txdb", "postgres",
		fmt.Sprintf("postgres://postgres@localhost/%v?sslmode=disable;user=%v;password=%v",
			cfg.DBName,
			cfg.DBUser,
			cfg.DBPassword))
}

func TestInsertUser(t *testing.T) {
	t.Run("InsertUser test", func(t *testing.T) {
		usr := database.Credentials{Username: "hello", Password: "worldherebro"}
		db, err := database.CreateDB("psql_txdb", cfg)
		if err != nil {
			t.Error(err)
		}
		err = db.InsertUser(usr)
		if err != nil {
			t.Error(err)
		}
		err = db.CheckUser(usr)
		if err != nil {
			t.Error(err)
		}
	})
}
