dbutil

abstraction for differences between mysql and sqlite for various utilities

also an implementation of a migration pattern for SQL powered applications

## Example
```go
package main

import (
    "fmt"
    "github.com/boourns/sqlgen"
)

type User struct {
    ID   int
    Name string
}

func main() {
    fmt.Printf("Create Table: %s", sqlgen.CreateTable(User{}))      
}
```
