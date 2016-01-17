package dbutil

import (
	"bytes"
	"reflect"
)

func CreateTable(i interface{}) string {
	t := reflect.TypeOf(i)

	if t.Kind() != reflect.Struct {
		panic("input must be of type struct")
	}

	buf := &bytes.Buffer{}
	err := createTable(buf, t)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
