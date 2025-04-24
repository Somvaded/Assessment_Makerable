package db


import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDatabase(DBurl string) *sql.DB {
    var err error
	dsn := DBurl
    DB, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    fmt.Println("Database connected!")
    return DB
}