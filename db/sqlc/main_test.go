package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/insta-app/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalln("cannot load config file:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to the db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
