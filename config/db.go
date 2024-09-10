package config

import (
	"database/sql"
	"fmt"
	"log"

	"expenze-io.com/schema/tables"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Ensure the connection string is correct
	connStr := "postgres://postgres:newpassword@localhost:5432/expenze-io?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Panic("Something went wrong in database: ", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

  tables.CreateTables(DB)

	fmt.Println("Database is up and running")
}
