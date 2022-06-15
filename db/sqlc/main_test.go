package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/souravgopal25/BankApplication/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err1 := util.LoadConfig("../..")
	if err1 != nil {
		log.Fatalln("Cannot Load config file", err1)
	}
	var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("Cannot create to db: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
