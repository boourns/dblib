package migrations

import (
	"database/sql"
	"errors"
	"github.com/boourns/dblib"
	"os"
	"testing"
)

var testEngine dblib.Engine

func clearMigrations() {
	testEngine.Exec("DELETE from migrations;")
}

func init() {
	url := os.Getenv("TEST_DATABASE_URL")

	if url == "" {
		panic("TEST_DATABASE_URL is not set, giving up")
	}

	testEngine = dblib.Connect(url)
}

func TestMigrateInit(t *testing.T) {
	for i := 0; i < 2; i++ {
		err := Init(testEngine)
		if err != nil {
			t.Errorf("Error running migrate.Init() %dth time: %v", i, err)
		}
	}
}

func TestMigrateRunsMigrationsOnce(t *testing.T) {
	Init(testEngine)
	clearMigrations()

	runCount := 0
	Run(1, "createUsers", func(tx *sql.Tx) error {
		runCount += 1
		return nil
	})
	if runCount != 1 {
		t.Errorf("Migration didn't run")
	}
	Run(1, "createUsers", func(tx *sql.Tx) error {
		t.Errorf("Migration 1 ran twice")
		return nil
	})
}

func TestMigrationRollsBackIfErrorReturned(t *testing.T) {
	Init(testEngine)
	clearMigrations()

	Run(1, "createUsers", func(tx *sql.Tx) error {
		return errors.New("something bad happened")
	})

	runCount := 0
	Run(1, "createUsers", func(tx *sql.Tx) error {
		runCount += 1
		return nil
	})

	if runCount != 1 {
		t.Errorf("failed migration blocked second run")
	}
}
