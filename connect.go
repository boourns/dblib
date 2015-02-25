package dbutil

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect(url string) *sql.DB {

	if url == "" {
		panic("DATABASE_URL is not set, expected  user:password@tcp(127.0.0.1:3306)/dbname")
	}

	db, err := sql.Open("mysql", url)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	return db
}
