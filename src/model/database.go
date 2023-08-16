package model

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	Driver string
	Source string
	Db     *sql.DB
}

func (d *Database) Connection() {
	db, err := sql.Open(d.Driver, d.Source) // "postgres", os.Getenv("POSTGRESQL_URL")

	if err != nil {
		log.Fatalln("[DATABASE]", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("[PING]", err)
	}

	d.Db = db
}

func (d *Database) Close() {
	d.Db.Close()
}
