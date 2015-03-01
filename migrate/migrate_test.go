package migrate

import (
	"github.com/boourns/dbutil"
	"os"
	"testing"
)

func init() {
	url := os.Getenv("TEST_DATABASE_URL")

	if url == "" {
		panic("TEST_DATABASE_URL is not set, giving up")
	}

	db = dbutil.Connect(url)
}

func TestMigrateRunsMigrationOnce(t *testing.T) {

}
