package db;

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var schema = `
CREATE TABLE jobs (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	url TEXT NOT NULL,
	description_url TEXT,
	tags TEXT,
	done Bool DEFAULT FALSE NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
`
var db *sqlx.DB

func GetDB() *sqlx.DB {
	if db != nil {
		return db
	}
	tmpDb, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=jobsearch sslmode=disable")
	log.Println("Connecting to DB ... ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	db = tmpDb
	return tmpDb
}

func Install() {
	log.Println("Installing DB")
	GetDB().MustExec(schema)
	log.Println("Done!")
}

