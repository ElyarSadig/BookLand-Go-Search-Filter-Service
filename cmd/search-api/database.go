package main

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var counts int64

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	// dsn := "host=localhost port=5432 user=admin password=password dbname=BookLandDB sslmode=disable timezone=UTC connect_timeout=5"

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
