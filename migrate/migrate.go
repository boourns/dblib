package migrate

import (
	"errors"
	"fmt"
	"github.com/boourns/dbutil"
)

var engine dbutil.Engine

func Run(version int, name string, migration func() error) bool {
	if migrationCompleted(version) == false {
		migration()
		insertMigration(version)
		return true
	}
	return false
}

func Init(e dbutil.Engine) error {
	engine = e

	if !e.TableExists("migrations") {
		sql := `CREATE TABLE migrations ( version INTEGER )`
		execute(sql)
	}

	if !e.TableExists("migrations") {
		return errors.New("could not create migrations table")
	}
	return nil
}

func migrationCompleted(version int) bool {
	rows, err := engine.Query("select 1 from migrations where version = ?", version)
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
	stmt, err := engine.Prepare(cmd)
	if err != nil {
		panic(fmt.Sprintf("Error preparing %s:%s", cmd, err))
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(fmt.Sprintf("Error executing %s:%s", cmd, err))
	}
}
