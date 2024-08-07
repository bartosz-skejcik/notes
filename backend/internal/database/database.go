package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
    var err error
    connStr := "user=myuser dbname=mydb sslmode=disable password=mypassword host=localhost port=5432"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }
    log.Println("Database connection established")
}