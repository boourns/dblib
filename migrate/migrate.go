package migrate

import (
	"database/sql"
	"errors"
	"fmt"
)

var db *sql.DB

func tableExists(name string) bool {
	var table, create string
	// try in mysql syntax
	err := db.QueryRow(`SHOW CREATE TABLE ?`, name).Scan(&table, &create)
	if err == nil {
		return true
	}

	// sqlite3 syntax
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name=?;`, name).Scan(&table)
	if table == "migrations" && err == nil {
		return true
	}

	return false
}

func Run(version int, name string, migration func() error) bool {
	if migrationCompleted(version) == false {
		migration()
		insertMigration(version)
		return true
	}
	return false
}

func Init(d *sql.DB) error {
	db = d

	if !tableExists("migrations") {
		sql := `CREATE TABLE migrations ( version INTEGER )`
		execute(sql)
	}

	if !tableExists("migrations") {
		return errors.New("could not create migrations table")
	}
	return nil
}

func migrationCompleted(version int) bool {
	rows, err := db.Query("select 1 from migrations where version = ?", version)
	if err != nil {
		panic(fmt.Sprintf("Error checking migration: %s", err))
	}
	defer rows.Close()
	return rows.Next()
}

func insertMigration(version int) {
	execute(fmt.Sprintf("INSERT INTO migrations(version) VALUES(%d)", version))
}

func execute(cmd string) {
	stmt, err := db.Prepare(cmd)
	if err != nil {
		panic(fmt.Sprintf("Error preparing %s:%s", cmd, err))
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(fmt.Sprintf("Error executing %s:%s", cmd, err))
	}
}
