# gorusgender [![gorusgender](https://github.com/ythosa/gorusgender/actions/workflows/go.yml/badge.svg)](https://github.com/ythosa/gorusgender/actions/workflows/go.yml) [![GoDoc](https://godoc.org/github.com/ythosa/gorusgender?status.svg)](https://pkg.go.dev/github.com/ythosa/gorusgender)

Go library for determining gender by last name, first name or middle name in Russian

**Example:**
```go
package main

import (
	"fmt"

	"github.com/ythosa/gorusgender"
)

func main() {
	gender := gorusgender.GetGender("Екатерина", "Сергеевна", "Иванова")
	fmt.Println(gender) // female
}
```
