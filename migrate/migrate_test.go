package migrate

import (
	"github.com/boourns/dbutil"
	"os"
	"testing"
)

var testEngine dbutil.Engine

func clearMigrations() {
	testEngine.Exec()
}

func init() {
	url := os.Getenv("TEST_DATABASE_URL")

	if url == "" {
		panic("TEST_DATABASE_URL is not set, giving up")
	}

	testEngine = dbutil.Connect(url)
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

}
