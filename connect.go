package dbutil

import (
	"database/sql"
	"log"

	"net/url"

	_ "github.com/go-sql-driver/mysql" // mysql adapter
	_ "github.com/mattn/go-sqlite3"    // sqlite3 adapter
)

// Connect accepts a database url like 'mysql://user:password@tcp(127.0.0.1:3306)/dbname' or 'sqlite3://data.db'
func Connect(dburl string) *sql.DB {
	parsed, err := url.Parse(dburl)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(parsed.Scheme, parsed.Path)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	return db
}
