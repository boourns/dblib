

package migrations

import (
	"database/sql"
	"fmt"
	"github.com/boourns/dblib"
)

func sqlFieldsForMigrations() string {
  return "Migrations.ID,Migrations.MigrationID,Migrations.CreatedAt" // ADD FIELD HERE
}

func loadMigrations(rows *sql.Rows) (*Migrations, error) {
	ret := Migrations{}

	err := rows.Scan(&ret.ID,&ret.MigrationID,&ret.CreatedAt) // ADD FIELD HERE
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func SelectMigrations(tx dblib.DBLike, cond string, condFields ...interface{}) ([]*Migrations, error) {
  ret := []*Migrations{}
  sql := fmt.Sprintf("SELECT %s from Migrations %s", sqlFieldsForMigrations(), cond)
	rows, err := tx.Query(sql, condFields...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
    item, err := loadMigrations(rows)
    if err != nil {
      return nil, err
    }
    ret = append(ret, item)
  }
  rows.Close()
  return ret, nil
}

func (s *Migrations) Update(tx dblib.DBLike) error {
		stmt, err := tx.Prepare(fmt.Sprintf("UPDATE Migrations SET ID=?,MigrationID=?,CreatedAt=? WHERE Migrations.ID = ?", )) // ADD FIELD HERE

		if err != nil {
			return err
		}

    params := []interface{}{s.ID,s.MigrationID,s.CreatedAt} // ADD FIELD HERE
    params = append(params, s.ID)

		_, err = stmt.Exec(params...)
		if err != nil {
			return err
		}

    return nil
}

func (s *Migrations) Insert(tx dblib.DBLike) error {
		stmt, err := tx.Prepare("INSERT INTO Migrations(MigrationID,CreatedAt) VALUES(?,?)") // ADD FIELD HERE
		if err != nil {
			return err
		}

		result, err := stmt.Exec(s.MigrationID,s.CreatedAt) // ADD FIELD HERE
		if err != nil {
			return err
    }

    s.ID, err = result.LastInsertId()
    if err != nil {
      return err
    }
	  return nil
}

func (s *Migrations) Delete(tx dblib.DBLike) error {
		stmt, err := tx.Prepare("DELETE FROM Migrations WHERE ID = ?")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(s.ID)
		if err != nil {
			return err
    }

	  return nil
}

func CreateMigrationsTable(tx dblib.DBLike) error {
		stmt, err := tx.Prepare(`



CREATE TABLE IF NOT EXISTS Migrations (
  
    ID INTEGER PRIMARY KEY,
   
    CreatedAt "DATETIME"
  
);

`)
		if err != nil {
			return err
		}

		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	  return nil
}
