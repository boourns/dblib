package migrations

import (
	"database/sql"
	"time"

	"github.com/boourns/dblib"
)

type Migrations struct {
	ID          int64
	MigrationID int64
	CreatedAt   time.Time `sqlType:"DATETIME"`
}

func DefineMigration(db *sql.DB, id int64, migration func(dblib.Queryable) error) error {
	existing, err := SelectMigrations(db, "WHERE MigrationID = ?", id)
	if err != nil {
		// Old documents are missing a MigrationID column, so try to add it seamlessly :( :)
		_, err2 := db.Exec("ALTER TABLE Migrations ADD COLUMN MigrationID INTEGER")

		if err2 != nil {
			return err2
		}
		_, err2 = db.Exec("UPDATE Migrations SET MigrationID=ID")
		if err2 != nil {
			return err2
		}

		// Try original query again
		existing, err = SelectMigrations(db, "WHERE MigrationID = ?", id)
		if err != nil {
			panic(err)
		}
	}

	if len(existing) != 0 {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = migration(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	completed := Migrations{MigrationID: id, CreatedAt: time.Now()}
	err = completed.Insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
