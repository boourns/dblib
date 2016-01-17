package dbutil

import (
	"log"
	"testing"
)

type User struct {
	ID   int
	Name string
}

func TestCreateTable(t *testing.T) {
	log.Printf(CreateTable(User{}))
}
