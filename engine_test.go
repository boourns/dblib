package dbutil

import (
	"log"
	"os"
	"testing"
)

var engine Engine

func init() {
	url := os.Getenv("TEST_DATABASE_URL")

	if url == "" {
		panic("TEST_DATABASE_URL is not set, giving up")
	}

	log.Println("Connnecting to ", url)

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
		return
	}

	tables, err := engine.Tables()
	if err != nil {
		t.Errorf("Error listing tables: %v\n", err)
		return
	}

	found := false
	for _, v := range tables {
		if v == "User" {
			found = true
		}
	}
	if found == false {
		t.Errorf("New table wasn't in list of tables, list is %v", tables)
		return
	}

	exist := engine.TableExists("User")
	if !exist || err != nil {
		t.Errorf("TableExists failed: exist %v err %v\n", exist, err)
		return
	}

	result, err := engine.Exec("INSERT INTO User(Name) VALUES (?)", "tom")
	if err != nil {
		t.Errorf("Insert failed: result %v err %v\n", result, err)
		return
	}

	if num, err := result.RowsAffected(); num != 1 || err != nil {
		t.Errorf("Insert into new table failed, num %v err %v", num, err)
		return
	}
}
