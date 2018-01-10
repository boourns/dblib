package dbutil

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql" // mysql adapter
	_ "github.com/mattn/go-sqlite3"    // sqlite3 adapter

	"github.com/boourns/dbutil/engine/mysql"
	"github.com/boourns/dbutil/engine/sqlite3"
)

// Connect accepts a database url like 'mysql://user:password@tcp(127.0.0.1:3306)/dbname' or 'sqlite3://data.db'
func Connect(dburl string) Engine {
	parts := strings.Split(dburl, "://")

	if len(parts) != 2 {
		log.Fatalf("Failed to connect: expected %s to be shaped like mysql://user:password@tcp(127.0.0.1:3306)/dbname)")
	}

	db, err := sql.Open(parts[0], parts[1])
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to %s database %s: %s", parts[0], parts[1], err)
	}

	engine := engineForScheme(parts[0], db)
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
