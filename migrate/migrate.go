package migrate

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

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

	var table, create string
	err := db.QueryRow(`SHOW CREATE TABLE migrations`).Scan(&table, &create)
	if err != nil {
		sql := `CREATE TABLE migrations ( version INTEGER )`
		execute(sql)
	}

	err = db.QueryRow(`SHOW CREATE TABLE migrations`).Scan(&table, &create)
	if err != nil {
		return err
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
