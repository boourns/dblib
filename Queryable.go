package dblib

import (
	"database/sql"
)

//DBLike is an interface to unify *sql.TX and *sql.DB for functions
//Deprecated: DBLike has been renamed to Queryable
type DBLike Queryable

//Queryable is an interface to unify *sql.TX and *sql.DB for functions.
type Queryable interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	Prepare(string) (*sql.Stmt, error)
	Exec(string, ...interface{}) (sql.Result, error)
}
