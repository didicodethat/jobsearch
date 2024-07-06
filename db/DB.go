package db;

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sqlx.DB

func GetDB() *sqlx.DB {
	if db != nil {
		return db
	}
	tmpDb, err := sqlx.Connect("sqlite3", "resources/db/jobsearch.sqlite")
	log.Print("Connecting to DB ... ")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected")
	db = tmpDb
	return tmpDb
}
