package dbutil

import (
	"os"
	"testing"
)

var engine Engine

func init() {
	url := os.Getenv("TEST_DATABASE_URL")

	if url == "" {
		panic("TEST_DATABASE_URL is not set, giving up")
	}

	engine = Connect(url)
}

func TestConnect(t *testing.T) {
	err := engine.Ping()
	if err != nil {
		t.Errorf("ping failed: %v", err)
	}
}

type User struct {
	ID   int
	Name string
}

func TestCreateTable(t *testing.T) {
	err := engine.CreateTable(User{})
	if err != nil {
		t.Errorf("Error creating table: %v\n", err)
	}

	tables, err := engine.Tables()
	if err != nil {
		t.Errorf("Error listing tables: %v\n", err)
	}

	if len(tables) != 1 || tables[0] != "User" {
		t.Errorf("New table not found, tables %v", tables)
	}

	result, err := engine.Exec("INSERT INTO User(Name) VALUES (?)", "tom")
	if num, err := result.RowsAffected(); num != 1 || err != nil {
		t.Errorf("Insert into new table failed, num %v err %v", num, err)
	}
}
