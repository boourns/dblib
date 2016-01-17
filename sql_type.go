package dbutil

import (
	"fmt"
	"log"
	"reflect"
)

func sqlType(f reflect.StructField) string {
	var t string
	switch f.Type.Name() {
	case "string":
		t = "VARCHAR(255)"
	case "int", "int64":
		t = "INTEGER"
	default:
		log.Fatalf("Unknown SQL type for go field %s", f.Type.Name())
	}

	if f.Name == "ID" || f.Name == "id" || f.Name == "Id" {
		t = fmt.Sprintf("%s PRIMARY KEY", t)
	}
	return t
}
