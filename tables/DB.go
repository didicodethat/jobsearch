package tables;

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
)

var db *sql.DB

func GetDB() *sql.DB {
	if db != nil {
		return db
	}
	tmpDb, err := sql.Open("sqlite3", "./db/jobsearch.sqlite")
	fmt.Print("Connecting to DB ... ")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDb
	return tmpDb
}
