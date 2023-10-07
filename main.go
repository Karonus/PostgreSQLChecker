package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")+"?sslmode=disable")
	if err != nil {
		os.Exit(-1)
		return
	}

	err = db.Ping()
	if err != nil {
		os.Exit(-1)
		return
	}

	os.Exit(0)
}
