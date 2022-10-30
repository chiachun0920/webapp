package main

import (
	"log"
	"os"
	"testing"
	"webapp/pkg/db"
)

var app application

func TestMain(m *testing.M) {
	app.Session = getSession()

	// setup db connection in test
	app.DSN = "host=postgres port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

	pathToTemplates = "./../../templates/"

	os.Exit(m.Run())
}
