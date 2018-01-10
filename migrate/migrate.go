package migrate

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/boourns/dbutil"
)

var engine dbutil.Engine

func Run(version int, name string, migration func(tx *sql.Tx) error) (bool, error) {
	complete, err := migrationCompleted(version)
	if err != nil {
		return false, err
	}

	if !complete {
		tx, err := engine.Begin()
		if err != nil {
			return false, err
		}

		err = migration(tx)
		if err != nil {
			tx.Rollback()
			return false, err
		}

		err = insertMigration(tx, version)
		if err != nil {
			tx.Rollback()
			return false, err
		}

		err = tx.Commit()
		if err != nil {
			return false, err
		}

		return true, nil
	}
	return false, nil
}

func Init(e dbutil.Engine) error {
	engine = e

	if !e.TableExists("migrations") {
		fmt.Printf("creating migrations table..\n")
		sql := `CREATE TABLE migrations ( version INTEGER )`
		_, err := e.Exec(sql)
		if err != nil {
			return err
		}
	}

	if !e.TableExists("migrations") {
		tables, _ := e.Tables()
		return errors.New(fmt.Sprintf("could not create migrations table, tables %v", tables))
	}
	return nil
}

func migrationCompleted(version int) (bool, error) {
	rows, err := engine.Query("select 1 from migrations where version = ?", version)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}

func insertMigration(tx *sql.Tx, version int) error {
	_, err := engine.Exec("INSERT INTO migrations(version) VALUES(?)", version)
	return err
}
