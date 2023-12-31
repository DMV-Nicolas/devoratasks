package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/DMV-Nicolas/DevoraTasks/util"
	_ "github.com/lib/pq"
)

var testQueries Store
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	testQueries = NewStore(testDB)

	os.Exit(m.Run())
}
