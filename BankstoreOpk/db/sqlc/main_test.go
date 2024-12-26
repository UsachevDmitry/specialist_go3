
package db

import (
	"context"
	"log"
	"os"
	"testing"
	"github.com/jackc/pgx/v5/pgxpool"

)

const (
	dbSource = "postgresql://postgres:P@ssw0rd@localhost:5432/bankdb?sslmode=disable"
)

var testDB *pgxpool.Pool
var testQueries *Queries

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("can not connect to db", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)

	os.Exit(m.Run())
}