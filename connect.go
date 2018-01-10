package dbutil

import (
	"database/sql"
	"fmt"
	"log"

	"net/url"

	_ "github.com/go-sql-driver/mysql" // mysql adapter
	_ "github.com/mattn/go-sqlite3"    // sqlite3 adapter

	"github.com/boourns/dbutil/engine/mysql"
	"github.com/boourns/dbutil/engine/sqlite3"
)

// Connect accepts a database url like 'mysql://user:password@tcp(127.0.0.1:3306)/dbname' or 'sqlite3://data.db'
func Connect(dburl string) Engine {
	parsed, err := url.Parse(dburl)
	if err != nil {
		log.Fatal(err)
	}

	hostname := parsed.Hostname()
	var path string
	switch {
	case hostname == "" && parsed.Path == "":
		// remain in memory
		path = ""
	case hostname == "" && parsed.Path != "":
		// like sqlite3:///path/to/file
		path = parsed.Path
	case hostname != "" && parsed.Path == "":
		// like sqlite3://test.db
		path = hostname
	case hostname != "" && parsed.Path != "":
		// like sqlite3://./test.db
		path = fmt.Sprintf("%s/%s", hostname, parsed.Path)
	}

	db, err := sql.Open(parsed.Scheme, path)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	engine := engineForScheme(parsed.Scheme, db)
	return engine
}

func engineForScheme(scheme string, db *sql.DB) Engine {
	switch scheme {
	case "sqlite3":
		return sqlite3.Engine{db}
	case "mysql":
		return mysql.Engine{db}
	}
	return nil
}
