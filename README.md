dblib

Utilities for interacting with SQL databases in Golang

## Features

`dblib.Queryable` - Interface to unify \*sql.TX and \*sql.DB for functions that should accept either.

`dblib.Transact` - Transaction wrapper

`github.com/boourns/dblib/migrations` - Migrations engine

`github.com/boourns/dblib/query` - Query builder
