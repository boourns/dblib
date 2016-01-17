sqlgen

generate SQL from golang static types

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
