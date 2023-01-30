package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:root@localhost:5432/gqlvf?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M){
	con, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database! ", err)
		return;
	}

	testQueries = New(con)

	os.Exit(m.Run())
}
