package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"expenze-io.com/internal/tables"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	connStr := os.Getenv("PG_CONNSTR")

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Panic("Something went wrong in database: ", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	tables.CreateTables(DB)

	fmt.Println("Database is up and running")
}
