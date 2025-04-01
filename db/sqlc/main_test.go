package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/insta-app/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalln("cannot load config file:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to the db:", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
