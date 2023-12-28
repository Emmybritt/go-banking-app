package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/emmybritt/bank_app/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB


func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Failed to log configurations", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}else {
		fmt.Println("Connection succesful")
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}