package db

import (
	"database/sql"
	"log"
	"os"
	"stock-prediction-backend/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource )
	if err != nil {
		log.Fatal("Cannot connec to db", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
