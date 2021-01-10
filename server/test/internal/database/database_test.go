package database_test

import (
	"fmt"
	"log"

	"github.com/DATA-DOG/go-txdb"
	"github.com/kembo91/kode-test-task/server/internal/database"
	"github.com/kembo91/kode-test-task/server/internal/utils"
	_ "github.com/lib/pq"
)

var cfg = utils.GetDBTestConfig()
var db = createTestingDb(cfg)

func createTestingDb(cfg utils.DBConfig) *database.Database {
	txdb.Register("psql_txdb", "postgres",
		fmt.Sprintf("postgres://postgres@localhost/%v?sslmode=disable;user=%v;password=%v",
			cfg.DBName,
			cfg.DBUser,
			cfg.DBPassword))
	db, err := database.CreateDB("psql_txdb", cfg)
	if err != nil {
		log.Fatal("cannot create test db instance")
	}
	return db
}
