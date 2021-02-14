package repo_test

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"
	"user/pkg/repo"

	_ "github.com/lib/pq"
)

var (
	testDB      *sql.DB
	testQueries *repo.Queries
)

// TODO: replace with config env
var dbSource = "postgres://postgres:secret@localhost:10521/user_service?sslmode=disable"
var dbDriver = "postgres"

func TestMain(m *testing.M) {
	var err error

	// get db conn
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("could not open database connection: %v", err)
	}

	// try connection
	for n := 0; n < 3; n++ {
		time.Sleep(3 * time.Second)
		if err = testDB.Ping(); err == nil {
			break
		}
	}
	if err != nil {
		log.Fatalf("database connection could not be established: %v", err)
	}

	// init testing repo.Queries
	testQueries = repo.New(testDB)

	os.Exit(m.Run())
}
